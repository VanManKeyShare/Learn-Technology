import time
import redis
from typing import cast
from typing import Optional
from fastapi import Request
from starlette.responses import JSONResponse
from starlette.middleware.base import BaseHTTPMiddleware


class RateLimitMiddleware(BaseHTTPMiddleware):
    """RATE LIMIT MIDDLEWARE USING REDIS LIMITS THE NUMBER OF REQUESTS PER IP ADDRESS"""

    def __init__(
        self,
        app,
        redis_client: redis.Redis,
        max_requests: int = 60,
        decay_seconds: int = 60,
        exclude_paths: Optional[list] = None,
    ):
        """
        INITIALIZE THE RATE LIMIT MIDDLEWARE
        Args:
            app: FASTAPI APPLICATION
            redis_client: REDIS CLIENT INSTANCE
            max_requests: MAXIMUM NUMBER OF REQUESTS ALLOWED (DEFAULT: 60)
            decay_seconds: TIME WINDOW IN SECONDS (DEFAULT: 60)
            exclude_paths: LIST OF PATHS TO EXCLUDE FROM RATE LIMITING (OPTIONAL)
        """
        super().__init__(app)
        self.redis_client = redis_client
        self.max_requests = max_requests
        self.decay_seconds = decay_seconds
        self.exclude_paths = exclude_paths or []

    def _get_client_ip(self, request: Request) -> str:
        # EXTRACT CLIENT IP FROM REQUEST
        # CHECK X-Forwarded-For HEADER FIRST (FOR PROXIES)
        if "x-forwarded-for" in request.headers:
            return request.headers["x-forwarded-for"].split(",")[0].strip()

        # FALLBACK TO CLIENT ADDRESS
        return request.client.host if request.client else "unknown"

    def _get_rate_limit_key(self, ip: str, path: str) -> str:
        # GENERATE REDIS KEY FOR RATE LIMITING
        return f"rate_limit:{ip}:{path}"

    async def dispatch(self, request: Request, call_next):
        # PROCESS THE REQUEST AND APPLY RATE LIMITING
        # SKIP RATE LIMITING FOR EXCLUDED PATHS
        if request.url.path in self.exclude_paths:
            return await call_next(request)

        # GET CLIENT IP
        client_ip = self._get_client_ip(request)

        # GENERATE RATE LIMIT KEY
        rate_limit_key = self._get_rate_limit_key(client_ip, request.url.path)

        try:
            # GET CURRENT REQUEST COUNT
            current_count = self.redis_client.get(rate_limit_key)

            if current_count is None:
                # FIRST REQUEST IN THIS WINDOW
                self.redis_client.setex(rate_limit_key, self.decay_seconds, 1)
                request.state.rate_limit_remaining = self.max_requests - 1
            else:
                current_count = int(cast(bytes, current_count))
                if current_count >= self.max_requests:
                    # RATE LIMIT EXCEEDED
                    ttl = self.redis_client.ttl(rate_limit_key)
                    ttl = int(cast(bytes, ttl))
                    return JSONResponse(
                        status_code=429,
                        content={
                            "status_code": 429,
                            "success": False,
                            "message": "TOO MANY REQUESTS",
                            "code": "RATE_LIMIT_EXCEEDED",
                            "data": {
                                "retry_after": ttl if ttl > 0 else self.decay_seconds,
                                "limit": self.max_requests,
                                "window": self.decay_seconds,
                            },
                        },
                        headers={
                            "Retry-After": str(ttl if ttl > 0 else self.decay_seconds),
                            "X-RateLimit-Limit": str(self.max_requests),
                            "X-RateLimit-Remaining": "0",
                            "X-RateLimit-Reset": str(
                                int(time.time())
                                + (ttl if ttl > 0 else self.decay_seconds)
                            ),
                        },
                    )

                # INCREMENT REQUEST COUNT
                self.redis_client.incr(rate_limit_key)
                request.state.rate_limit_remaining = (
                    self.max_requests - current_count - 1
                )

        except Exception as e:
            # LOG ERROR BUT ALLOW REQUEST TO PROCEED
            print(f"RATE LIMIT MIDDLEWARE ERROR: {str(e)}")
            return await call_next(request)

        # PROCESS THE REQUEST
        response = await call_next(request)

        # ADD RATE LIMIT HEADERS
        response.headers["X-RateLimit-Limit"] = str(self.max_requests)
        response.headers["X-RateLimit-Remaining"] = str(
            getattr(request.state, "rate_limit_remaining", self.max_requests)
        )

        ttl = self.redis_client.ttl(rate_limit_key)
        ttl = int(cast(bytes, ttl))
        response.headers["X-RateLimit-Reset"] = str(
            int(time.time()) + (ttl if ttl > 0 else self.decay_seconds)
        )

        return response
