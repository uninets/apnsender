package apnsender

import (
    "crypto/tls"
    "net"
    "encoding/hex"
    "encoding/binary"
    "bytes"
    "errors"
    "strconv"
)



type Device struct {
    Token string
    Timestamp uint32
}

type Message struct {
    Token string
    Content []byte
}

type sender interface {
    Connect() error
    Send(chan []byte, chan error)
    Feedback() []Device
}

type APS struct {
    Config APSConfig
    Connection net.Conn
}

func New(config APSConfig) (a APS, err error) {
    if checkConfig(config) {
        a = APS{Config: config}
    } else {
        err = errors.New("Config check failed.")
        return
    }

    return
}

func (a *APS) Connect() (err error) {
    // load certificate from files
    cert, err := tls.LoadX509KeyPair(
        a.Config.sslCert,
        a.Config.sslKey,
    )
    if err != nil { return }

    // connect to apple push service
    conn, err := net.Dial(
        "tcp",
        a.Config.host + ":" + strconv.Itoa(int(a.Config.port)),
    )

    if err != nil { return }

    // wrap socket with tls client
    tlsconn := tls.Client(conn, &tls.Config{
        Certificates: []tls.Certificate{cert},
    })

    // test connection
    err = tlsconn.Handshake()
    if err != nil { return }

    a.Connection = tlsconn

    return
}

func checkConfig(config APSConfig) bool {
    return true
}

func (a *APS) Send(c chan Message, errors chan error) {
    go func() {
        for {
            message := <- c
            payload, err := a.generatePayload(message)

            if err != nil {
                errors <- err
            } else {
                a.Connection.Write(payload)
            }
        }
    }()
}

func (a *APS) generatePayload(m Message) (pdu []byte, err error) {

    deviceToken, err := hex.DecodeString(m.Token)
    if err != nil { return }

    json_s := m.Content

    payload := bytes.NewBuffer([]byte{})
    // command
    binary.Write(payload, binary.BigEndian, uint8(1))
    // transaction id
    binary.Write(payload, binary.BigEndian, uint32(1))
    // expiration time, 1h
    binary.Write(payload, binary.BigEndian, uint32(3600))
    // device token
    binary.Write(payload, binary.BigEndian, uint16(len(deviceToken)))
    binary.Write(payload, binary.BigEndian, deviceToken)
    // push notification
    binary.Write(payload, binary.BigEndian, uint16(len(json_s)))
    binary.Write(payload, binary.BigEndian, json_s)

    // binary
    pdu = payload.Bytes()

    return
}
