 go_different_ways_to_post

https://www.getpostman.com/docs/requests
http://stackoverflow.com/questions/26723467/what-is-the-difference-between-form-data-x-www-form-urlencoded-and-raw-in-the-p

= 
__application/x-www-form-urlencoded :__   Transcode all characters before uploading (default).
__multipart/form-data :__   No transcoding. You must use this value when your form has file upload controls.
__text/plain :__    Convert spaces to "+", but no transcoding for special characters.

