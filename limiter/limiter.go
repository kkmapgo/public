package limiter

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(num int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: num,
		bucket:         make(chan int, num),
	}
}

func (cl *ConnLimiter) GetConnection() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("Connections have reach then rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConnection() {
	c := <-cl.bucket
	log.Printf("Connection had been released, %d \n", c)
}
