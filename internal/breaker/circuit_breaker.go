package breaker

import (
	"github.com/sony/gobreaker"
	"github.com/iaa-gamecenter/iaa-gamecenter-sdk-go/config"
	"sync"
	"time"
)

type ServiceCircuitBreaker struct {
}

var (
	breakerMap = make(map[string]*gobreaker.CircuitBreaker)
	mutex      sync.Mutex
)

func NewBreaker() *ServiceCircuitBreaker {
	return &ServiceCircuitBreaker{}
}

func (b *ServiceCircuitBreaker) GetBreaker(name string, config *config.BreakerConfig) *gobreaker.CircuitBreaker {
	key := name + config.GetKey()
	breaker := breakerMap[key]
	if breaker == nil {
		mutex.Lock()
		breaker = breakerMap[key]
		if breaker == nil {
			breaker = b.createBreaker(key, config)
			breakerMap[key] = breaker
		}
		mutex.Unlock()
	}
	return breaker
}

func (b *ServiceCircuitBreaker) createBreaker(name string, config *config.BreakerConfig) *gobreaker.CircuitBreaker {
	var st gobreaker.Settings
	st.Name = name
	st.MaxRequests = config.MaxRequests
	st.Timeout = time.Second * time.Duration(config.BreakerTimeout)
	st.Interval = time.Second * time.Duration(config.Interval)
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= config.TripRequest && failureRatio >= config.FailureRatio
	}
	return gobreaker.NewCircuitBreaker(st)
}
