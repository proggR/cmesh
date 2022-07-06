package cmtp
import "fmt"
func main() {
    fmt.Println("CMTP Invoked")
}
// 
// type CMTPMessage struct {
//   headers CMTPMessageHeaders
//   request CMTPRequest
//   response CMTPResponse
// }
//
// type CMBlob struct {
//
// }
//
// type JWT struct {
//
// }
//
// type CMTPRequest struct {
//   host string
//   path string
//   iamSessionToken JWT
//   payload CMBlob
// }
//
// type CMTPResponse struct {
//   status Numeric
//   host string
//   path string
//   iamSessionToken JWT
//   payload CMBlob
// }
//
//
// type CMTPMessageHeaders struct {
//   status Numeric
//   version string
//   transferred Numeric
//   referrerPolicy string
//   request CMTPRequestHeader
//   response CMTPResponseHeader
// }
//
// type CMTPRequestHeader struct {
//   accept string
//   acceptEncoding string
//   acceptLanguage string
//   connection string
//   host string
//   ifModSince string
//   ifNoneMatch string
//   referrer string
//   secFetchDest string
//   secFetchMode string
//   secFetchUser string
//   userAgent string
// }
//
// type CMTPResponseHeader struct {
//   acceptRanges string
//   cacheControl string
//   contentEncoding string
//   contentLength: Numeric
//   contentSecurityPolicy string
//   contentType string
//   date string
//   etag string
//   expires string
//   lastModified string
//   strictTransportSecurity string
//   vary string
// }
//
// func message_compiler(req CMTPRequest){
//   var header CMTPRequestHeader = load_request_header()
//   var msg CMTPMessage = request_wrap(header,req)
//   var res = response_await(msg)
//   return res
// }
//
// func load_message_header(){
//
// }
//
// func load_request_header(){
//
// }
//
// func response_await(msg CMTPMessage){
//
// }
//
// func request_wrap(header CMTPRequestHeader, request CMTPRequest){
//   var msg = CMTPMessage{load_message_header(header),request,nil}
// }
