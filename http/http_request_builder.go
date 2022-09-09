package http

// NewRequestBuilder 普通Builder工厂方法，新创建一个Request对象
// 工厂模式的实现其实有很多方法，比如抽象工厂 ，builder 等等，这里我们用builder来实现
type requestBuilder struct {
	req *Request
}

func NewRequestBuilder() *requestBuilder {
	return &requestBuilder{req: EmptyRequest()}
}

func (r *requestBuilder) AddMethod(method Method) *requestBuilder {
	r.req.method = method
	return r
}

func (r *requestBuilder) AddUri(uri Uri) *requestBuilder {
	r.req.uri = uri
	return r
}

func (r *requestBuilder) AddQueryParam(key, value string) *requestBuilder {
	r.req.queryParams[key] = value
	return r
}

func (r *requestBuilder) AddQueryParams(params map[string]string) *requestBuilder {
	for k, v := range params {
		r.req.queryParams[k] = v
	}
	return r
}

func (r *requestBuilder) AddHeader(key, value string) *requestBuilder {
	r.req.headers[key] = value
	return r
}
func (r *requestBuilder) AddHeaders(headers map[string]string) *requestBuilder {
	for k, v := range headers {
		r.req.headers[k] = v
	}
	return r
}

func (r *requestBuilder) AddBody(body interface{}) *requestBuilder {
	r.req.body = body
	return r
}

func (r *requestBuilder) Builder() *Request {
	return r.req
}
