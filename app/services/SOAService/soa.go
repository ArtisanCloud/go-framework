package SOAService

type SOA interface {
	SetHeader(obj interface{}) SOA
	SetBody(obj interface{}) SOA
	SetForm(obj interface{}) SOA

	BindJSON(obj interface{}) SOA
	BindHeader(obj interface{}) SOA
	Code(httpCode *int) SOA

	GET(url string) SOA
	POST(url string) SOA
	Do() error
}

type RspHeader struct {
	ContentType string `header:"Content-Type"`
}
