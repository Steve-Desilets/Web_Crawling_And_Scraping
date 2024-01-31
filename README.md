For the project included within this Github repository, we imagine that we are a data scientist at a company whose leadership team is interested in leveraging Go for its web crawling and scraping processes.  The leadership team at this company has had success in the past crawling across webpages and scraping data from them via Python's Scrapy module.  However, given that the leadership team believes that Go is faster at computing and processing than Python is, the leadership team has asked the data scientist to study whether Go has the same capability to crawl across webpages, scrape information, compile the data into a JSON format, and export the data into an output julia script file.  If so, then the leadership team may be interested in switching its data engineering code from leveraging Python to leveraging Go.

Accordingly, the data scientist in this study created the "web_crawling_and_scraping" Go script included within this repository to crawl across the Wikipedia pages for the following topics and scrape information from them.
    
    - Robotics
    - Robot
    - Reinforcement learning

For each of these webpages, the data scientist extracted the webpage's title, the webpage's URL, and the text from the body of the webpage.  The Go program then compiled this data into a JSON for each webpage and exported all the scraped data into the "scrapedOutputGo" output julia file. 

If I were the data scientist making a recommendation to the leadership team of this company based on the findings of this study, I likely would recommend that the company begin leveraging Go instead of Python for its data engineering projects.  That's because previous experiments at this company (as described in the Command_Line_Applications and Benchmarking-Project Github repositories) have proven that Go is a much faster language computationally than Python is.  Given that Go can also produce equally excellent outputs when crawling and scraping the internet, Go should be the data engineering language of choice for this company.
