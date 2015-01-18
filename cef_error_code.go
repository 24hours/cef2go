package chrome

type CefErrorCode int

const (
	ERR_NONE                            CefErrorCode = 0
	ERR_FAILED                          CefErrorCode = -2
	ERR_ABORTED                         CefErrorCode = -3
	ERR_INVALID_ARGUMENT                CefErrorCode = -4
	ERR_INVALID_HANDLE                  CefErrorCode = -5
	ERR_FILE_NOT_FOUND                  CefErrorCode = -6
	ERR_TIMED_OUT                       CefErrorCode = -7
	ERR_FILE_TOO_BIG                    CefErrorCode = -8
	ERR_UNEXPECTED                      CefErrorCode = -9
	ERR_ACCESS_DENIED                   CefErrorCode = -10
	ERR_NOT_IMPLEMENTED                 CefErrorCode = -11
	ERR_CONNECTION_CLOSED               CefErrorCode = -100
	ERR_CONNECTION_RESET                CefErrorCode = -101
	ERR_CONNECTION_REFUSED              CefErrorCode = -102
	ERR_CONNECTION_ABORTED              CefErrorCode = -103
	ERR_CONNECTION_FAILED               CefErrorCode = -104
	ERR_NAME_NOT_RESOLVED               CefErrorCode = -105
	ERR_INTERNET_DISCONNECTED           CefErrorCode = -106
	ERR_SSL_PROTOCOL_ERROR              CefErrorCode = -107
	ERR_ADDRESS_INVALID                 CefErrorCode = -108
	ERR_ADDRESS_UNREACHABLE             CefErrorCode = -109
	ERR_SSL_CLIENT_AUTH_CERT_NEEDED     CefErrorCode = -110
	ERR_TUNNEL_CONNECTION_FAILED        CefErrorCode = -111
	ERR_NO_SSL_VERSIONS_ENABLED         CefErrorCode = -112
	ERR_SSL_VERSION_OR_CIPHER_MISMATCH  CefErrorCode = -113
	ERR_SSL_RENEGOTIATION_REQUESTED     CefErrorCode = -114
	ERR_CERT_COMMON_NAME_INVALID        CefErrorCode = -200
	ERR_CERT_DATE_INVALID               CefErrorCode = -201
	ERR_CERT_AUTHORITY_INVALID          CefErrorCode = -202
	ERR_CERT_CONTAINS_ERRORS            CefErrorCode = -203
	ERR_CERT_NO_REVOCATION_MECHANISM    CefErrorCode = -204
	ERR_CERT_UNABLE_TO_CHECK_REVOCATION CefErrorCode = -205
	ERR_CERT_REVOKED                    CefErrorCode = -206
	ERR_CERT_INVALID                    CefErrorCode = -207
	ERR_CERT_END                        CefErrorCode = -208
	ERR_INVALID_URL                     CefErrorCode = -300
	ERR_DISALLOWED_URL_SCHEME           CefErrorCode = -301
	ERR_UNKNOWN_URL_SCHEME              CefErrorCode = -302
	ERR_TOO_MANY_REDIRECTS              CefErrorCode = -310
	ERR_UNSAFE_REDIRECT                 CefErrorCode = -311
	ERR_UNSAFE_PORT                     CefErrorCode = -312
	ERR_INVALID_RESPONSE                CefErrorCode = -320
	ERR_INVALID_CHUNKED_ENCODING        CefErrorCode = -321
	ERR_METHOD_NOT_SUPPORTED            CefErrorCode = -322
	ERR_UNEXPECTED_PROXY_AUTH           CefErrorCode = -323
	ERR_EMPTY_RESPONSE                  CefErrorCode = -324
	ERR_RESPONSE_HEADERS_TOO_BIG        CefErrorCode = -325
	ERR_CACHE_MISS                      CefErrorCode = -400
	ERR_INSECURE_RESPONSE               CefErrorCode = -501
)
