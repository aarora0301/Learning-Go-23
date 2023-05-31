package problems

// rate limit with token bucket algorithm

// create x tokens per min
// clear window
// generate x tokens

// getToken
// see if token is available
// else dicard request or wait for token to be available

// Implementing rate limiting using a naive token bucket filter algorithm.

//Imagine you have a bucket that gets filled with tokens at the rate of 1 token per second.
//	The bucket can hold a maximum of N tokens. Implement a thread-safe class that lets threads
//get a token when one is available. If no token is available, then the token-requesting threads
//should block.
