import os
import redis
from typing import Union
from typing import Any, Optional
from pydantic import BaseModel
from sqlalchemy import text
from sqlalchemy import create_engine
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
from fastapi.exceptions import RequestValidationError
from fastapi.middleware.cors import CORSMiddleware
from starlette.exceptions import HTTPException as StarletteHTTPException
from Middleware_RATE_LIMIT import RateLimitMiddleware


class Success_Response_Model(BaseModel):
    status_code: int = 200
    success: bool = True
    message: str = "OK"
    code: Optional[Any] = None
    data: Optional[Any] = None
    meta: Optional[Any] = None
    errors: Optional[Any] = None


class Error_Response_Model(BaseModel):
    status_code: int = 422
    success: bool = False
    message: str = "Error"
    code: Optional[Any] = None
    data: Optional[Any] = None
    meta: Optional[Any] = None
    errors: Optional[Any] = None


app = FastAPI(
    title="API SERVER",
    description="API SERVER",
    version="1.0",
    docs_url="/vmk-docs",
    redoc_url=None,
    openapi_url="/vmk-openapi.json",
)

r = redis.Redis(
    host=os.getenv("REDIS_HOST", "localhost"),
    port=int(os.getenv("REDIS_PORT", "6379")),
    password=os.getenv("REDIS_PASSWORD", ""),
    decode_responses=True,
)

# ADD CORS MIDDLEWARE
cors_origins = [
    "https://localhost",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=cors_origins,
    allow_credentials=True,
    allow_methods=["GET", "POST", "PUT", "DELETE", "PATCH"],
    allow_headers=["Content-Type", "Authorization"],
    expose_headers=["X-RateLimit-Limit", "X-RateLimit-Remaining", "X-RateLimit-Reset"],
    max_age=3600,
)

# ADD RATE LIMIT MIDDLEWARE
app.add_middleware(
    RateLimitMiddleware,
    redis_client=r,
    max_requests=60,
    decay_seconds=60,
    exclude_paths=["/vmk-docs", "/vmk-openapi.json"],  # EXCLUDE DOCS ENDPOINTS
)


@app.get(
    "/",
    tags=["Main"],
    summary="HOME",
    description="RETURNS A WELCOME MESSAGE AND REDIS COUNT.",
    response_model=Success_Response_Model,
    responses={200: {"model": Success_Response_Model}},
)
def root():
    r.incr("Counter")
    return JSONResponse(
        status_code=200,
        content=Success_Response_Model(
            status_code=200,
            success=True,
            message="OK",
            data={
                "Title": "FASTAPI + MARIADB + REDIS",
                "Request Counter": r.get("Counter"),
                "Request Rate Limiting": "60 REQUESTS PER MINUTE",
            },
        ).model_dump(),
    )


@app.get(
    "/db",
    tags=["Main"],
    summary="CHECK DB CONNECTION",
    description="CHECKS THE CONNECTION TO THE MARIADB DATABASE.",
    response_model=Union[Success_Response_Model, Error_Response_Model],
    responses={
        200: {"model": Success_Response_Model},
        500: {"model": Error_Response_Model},
    },
)
def check_db_connection():
    user = os.getenv("MYSQL_USER")
    password = os.getenv("MYSQL_PASSWORD")
    host = os.getenv("DB_HOST")
    port = os.getenv("DB_PORT")
    database = os.getenv("MYSQL_DATABASE")

    DB_URL = f"mysql+pymysql://{user}:{password}@{host}:{port}/{database}"

    engine = create_engine(DB_URL, echo=False)
    try:
        with engine.connect() as conn:
            version = conn.execute(text("SELECT VERSION()")).scalar()
            return JSONResponse(
                status_code=200,
                content=Success_Response_Model(
                    status_code=200,
                    success=True,
                    message="OK",
                    data={"message": "CONNECTED TO MARIADB", "version": version},
                ).model_dump(),
            )
    except Exception as e:
        return JSONResponse(
            status_code=500,
            content=Error_Response_Model(
                status_code=500,
                success=False,
                message="ERROR CONNECTING TO MARIADB",
                errors=str(e),
            ).model_dump(),
        )


@app.get(
    "/items/{item_id}",
    tags=["Main"],
    summary="GET AN ITEM BY ID",
    description="RETRIEVE AN ITEM USING ITS UNIQUE IDENTIFIER.",
    response_model=Union[Success_Response_Model, Error_Response_Model],
    responses={
        200: {"model": Success_Response_Model},
        404: {"model": Error_Response_Model},
        422: {"model": Error_Response_Model},
    },
)
def read_item(item_id: int, query: str | None = None):
    if item_id < 1000:
        return JSONResponse(
            status_code=422,
            content=Error_Response_Model(
                status_code=422,
                success=False,
                message="ITEM ID MUST BE GREATER THAN OR EQUAL TO 1000.",
            ).model_dump(),
        )

    if item_id > 2000:
        return JSONResponse(
            status_code=404,
            content=Error_Response_Model(
                status_code=404,
                success=False,
                message="KHÔNG TÌM THẤY ITEM VỚI ID NÀY.",
            ).model_dump(),
        )

    return JSONResponse(
        status_code=200,
        content=Success_Response_Model(
            status_code=200,
            success=True,
            message="OK",
            data={"item_id": item_id, "query": query},
        ).model_dump(),
    )


@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request: Request, exc: RequestValidationError):
    return JSONResponse(
        status_code=422,
        content=Error_Response_Model(
            status_code=422,
            success=False,
            message="INVALID DATA FORMAT",
            errors=exc.errors(),
        ).model_dump(),
    )


@app.exception_handler(StarletteHTTPException)
async def custom_http_exception_handler(request: Request, exc: StarletteHTTPException):
    if exc.status_code == 404:
        return JSONResponse(
            status_code=exc.status_code,
            content=Error_Response_Model(
                status_code=exc.status_code,
                success=False,
                message="API KHÔNG TỒN TẠI",
                errors=exc.detail,
            ).model_dump(),
        )

    # FALLBACK CHO LỖI KHÁC
    return JSONResponse(
        status_code=exc.status_code,
        content=Error_Response_Model(
            status_code=exc.status_code,
            success=False,
            message="KHÔNG XÁC ĐỊNH",
            errors=exc.detail,
        ).model_dump(),
    )
