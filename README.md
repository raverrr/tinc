# tinc
Script to aid in timing based hacks that require an interger is to be incremented on each request. By default tinc multiplies the suplied integer by 2 on each request. This was created originally as a one off for a POC using BENCHMARK() in a SQL injection.

# Usage:
 * -i int
 *   	--> Interger to increment by in each request 
    	
 * -n int
 *   	--> Number of requests to send 
    	
 * -u string
 *   	--> URL with 'ZCZC' where the interger to increment should be placed 
