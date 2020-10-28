## Steps to build and run the CLI tool
1. Use the Command: *go build filename.go* to build the file (Make sure that you are in the correct directory)    
2. Use the Command: *go run filename.go* to run the program    
3. There are two commands defined in this tool    
	i. **runTool**: Provides a JSON Object containing links    
		There are two flags defined for command runTool

		a. --url value     
		Example: cli --url cloudflare.com  (go run not required after build)    

		b. --profile value     
		Example: cli --profile 5 (go run not required after build)     

	*Both the flags can be used together in the runTool command*     
	*Example: cli runTool --profile 3 --url software_assessment.yasharavindra-wrangler.workers.dev*

	ii. **help, h**: Shows a list of commands or help for one command     

## Network Statistics Comparison    

![Network Statistics](https://raw.githubusercontent.com/yasharavindra18/Go-CLI/master/NetworkStatistics.PNG)    

In terms of network performance, latency of google.com was comparatively better than other domains with latency of **17.1971 ms.**     
Our latency is **18.5115 ms**, for the workers.dev sub-domain: software_assessment.yasharavindra-wrangler.workers.dev, which is better than the latency of twitter.com (30.4018 ms) and facebook.com (20.7985 ms).      
Here, the latency is measured in terms of round trip time.
