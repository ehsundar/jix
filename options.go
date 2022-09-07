package jix

func (j *Jixer[Req, Resp]) WithFillRequestFromHeader() *Jixer[Req, Resp] {
	return j.WithRequestExtractors(HeaderExtractor[Req])
}

func (j *Jixer[Req, Resp]) WithFillHeadersFromResponse(fill bool) *Jixer[Req, Resp] {
	j.fillResponseHeaders = fill
	return j
}

func (j *Jixer[Req, Resp]) WithFillRequestFromQuery() *Jixer[Req, Resp] {
	return j.WithRequestExtractors(QueryExtractor[Req])
}

func (j *Jixer[Req, Resp]) WithErrorToStatusMapping(m map[error]int) *Jixer[Req, Resp] {
	for err, code := range m {
		j.statusMapper[err] = code
	}
	return j
}

func (j *Jixer[Req, Resp]) WithRequestExtractors(ext ...RequestExtractor[Req]) *Jixer[Req, Resp] {
	j.requestExtractors = append(j.requestExtractors, ext...)
	return j
}
