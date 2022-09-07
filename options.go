package jix

func (j *Jixer[Req, Resp]) WithFillRequestFromHeader(fill bool) *Jixer[Req, Resp] {
	j.fillRequestHeaders = fill
	return j
}

func (j *Jixer[Req, Resp]) WithFillHeadersFromResponse(fill bool) *Jixer[Req, Resp] {
	j.fillResponseHeaders = fill
	return j
}

func (j *Jixer[Req, Resp]) WithFillRequestFromQuery(fill bool) *Jixer[Req, Resp] {
	j.fillQueries = fill
	return j
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
