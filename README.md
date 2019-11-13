This project aims to reproduce Golang's http.Client's issues with **keep-alive** and **TIME_WAIT**.
It also provides the solution to fix the reproduced problems.

To make things understandable, the project built as a story line going through each problem, 
its solution.

Order:

+ **default_noclose**  
Why it is important to close the Response.Body even if the body is not used.
+ **default_noread**  
Why it is important to read Response.Body even if the body is not used.
+ **clean**  
A clean solution as long as the simultaneous requests do not exceed http.DefaultClient's MaxIdleConnsPerHost.
+ **clean_fail**  
Showcases what happens when the simultaneous requests exceeds http.DefaultClient's MaxIdleConnsPerHost.
+ **cleaner**  
Configuring MaxIdleConnsPerHost to avoid temporary connection creation.
+ **cleaner_fail**  
Temporary connections are still created as long as the number of simultaenous connection exceeds configured
MaxIdleConnsPerHost.
+ **cleanest**  
Put a max cap on how any out standing connections allowed using MaxConnsPerHost.