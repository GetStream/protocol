import httpx
from stream_response import StreamResponse
from generic import T
from urllib.parse import quote

def build_path(path: str, path_params: dict) -> str:
    for k, v in path_params.items():
        path_params[k] = quote(
            v, safe=""
        )  # in case of special characters in the path. Known cases: chat message ids.

    return path.format(**path_params)



class BaseClient:
    client: httpx.Client
    basePath: str = "http://localhost:8080"

    def __init__(self):
        self.client = httpx.Client()

    def request(self, method: str, path: str, body=None, query_params:dict=None, path_params: dict = None, **kwargs) -> StreamResponse[T]:
       
        return self.client.request(method, self.basePath + build_path(path, path_params), params=query_params, json=body, **kwargs)
    
    
