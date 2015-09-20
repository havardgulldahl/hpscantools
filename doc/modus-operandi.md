# Modus Operandi#

A pseudo-code description on what we do, based on @matthew1001 code https://github.com/matthew1001/mrw-code/blob/master/opt/hp-scanner-monitor/monitor-scanner-for-matt.sh

1. Submit an initial request just to check the printer's online
    xml=$(curl -s -X GET http://$printerIP/WalkupScan/WalkupScanDestinations)

2. Send a POST request containing XML which describes us

    response=$(curl -s -v -X POST -d @ourdetails-matt.xml --header 'Content-Type: text/xml' http://$printerIP/WalkupScan/WalkupScanDestinations 2>&1 | grep Location)

3. Get `Location` header from HTTP POST

4. Send a GET request to check for new scan events

    xml=$(curl -s -X GET http://$printerIP/EventMgmt/EventTable?timeout=1200)

    if [[ "$xml" == *"PoweringDownEvent"* ]]; then

        # The printer is powering down

    if [[ "$xml" == *"$url"* ]]; then

        echo "XML response contains a new event for us"
        
        # Send a request to our unique URL to get any specific details we need
        xml=$(curl -s -X GET $url)
        if [[ -z "$xml" ]]; then
        # We didn't get a good response - we might have been disconnected

        # Send a request to get the scan status, abort if we don't get it
        xml=$(curl -s -X GET http://$printerIP/Scan/Status)

        response=$(curl -s -v -X POST -d @scandetails.xml --header 'Content-Type: text/xml' http://$printerIP/Scan/Jobs 2>&1 | grep Location)
        joburl = Location header

        # Send a request to the job 
        binaryurl=$(curl -s -X GET $joburl | xpath -q -e "/j:Job/ScanJob/PreScanPage/BinaryURL/text()")

        while true:
            # Send a request to get the current job state 
            jobstate=$(curl -s -X GET $joburl | xpath -q -e "/j:Job/j:JobState/text()")
            sleep 60
            #Get the image which is probably done by now"
            xml=$(curl -s -X GET -o /tmp/image.jpg http://$printerIP$binaryurl)
            break
     


