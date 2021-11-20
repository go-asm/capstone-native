package capstone

func (e *Engine) SkipData() *CsOptSkipdata {
	if e.skipData == nil {
		return nil
	}

	return e.skipData
}
