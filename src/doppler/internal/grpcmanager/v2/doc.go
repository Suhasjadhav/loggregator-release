package v2

import plumbing "plumbing/v2"

//go:generate hel
type BatcherSenderServer interface {
	plumbing.DopplerIngress_BatchSenderServer
}
