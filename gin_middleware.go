type xxWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w xxWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w xxWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ww := &xxWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = ww

		c.Next()
	}
}
