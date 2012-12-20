package apnsender

const (
    DefaultAPSHost      = "gateway.push.apple.com"
    DefaultAPSPort      = 2195
    DefaultFeedbackHost = "feedback.push.apple.com"
    DefaultFeedbackPort = 2196
    DefaultSSLCert      = "cert.pem"
    DefaultSSLKey       = "key.pem"
)

type APSConfig struct {
    host string
    port uint16
    feedbackHost string
    feedbackPort uint16
    sslCert string
    sslKey string
}

func DefaultConfig() *APSConfig {
    return &APSConfig{
        DefaultAPSHost,
        DefaultAPSPort,
        DefaultFeedbackHost,
        DefaultFeedbackPort,
        DefaultSSLCert,
        DefaultSSLKey,
    }
}

func (c *APSConfig) Host(host string) *APSConfig {
    c.host = host
    return c
}

func (c *APSConfig) Port(port int) *APSConfig {
    c.port = uint16(port)
    return c
}

func (c *APSConfig) FeedbackHost(host string) *APSConfig {
    c.feedbackHost = host
    return c
}

func (c *APSConfig) FeedbackPort(port int) *APSConfig {
    c.feedbackPort = uint16(port)
    return c
}

func (c *APSConfig) SSLCert(cert string) *APSConfig {
    c.sslCert = cert
    return c
}

func (c *APSConfig) SSLKey(key string) *APSConfig {
    c.sslKey = key
    return c
}

