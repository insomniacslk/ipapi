# ipapi

The package `ipapi` implements a small wrapper around the `ip-api.com` API.

There is only one method, `Get`, that receives an IP address and a proxy URL.

If the IP address is `nil`, ip-api.com will use the IP address that is seen on
their end.

If the proxy is not `nil`, the request will go through a proxy.
