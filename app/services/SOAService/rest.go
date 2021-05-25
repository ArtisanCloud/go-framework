package SOAService

import (
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
)

type RestService struct {
	url      string
	rsHeader interface{}
	rsBody   interface{}
	df       *dataflow.DataFlow
}

func NewRestService() *RestService {
	return &RestService{
	}
}

func (srv RestService) SetHeader(obj interface{}) SOA {
	srv.df.SetHeader(obj)
	return srv
}
func (srv RestService) SetBody(obj interface{}) SOA {
	srv.df.SetBody(obj)
	return srv
}

func (srv RestService) SetForm(obj interface{}) SOA {
	srv.df.SetForm(obj)
	return srv
}

func (srv RestService) BindJSON(obj interface{}) SOA {
	srv.df.BindJSON(obj)
	return srv
}
func (srv RestService) BindHeader(obj interface{}) SOA {
	srv.df.BindHeader(obj)
	return srv
}

func (srv RestService) Code(httpCode *int) SOA {
	srv.df.Code(httpCode)
	return srv
}

func (srv RestService) GET(url string) SOA {
	srv.df = gout.GET(url)

	return srv
}

func (srv RestService) POST(url string) SOA {
	srv.df = gout.POST(url)

	return srv
}

func (srv RestService) Do() error {
	err := srv.df.Do()

	return err
}
