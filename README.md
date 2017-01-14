Uses EXIF img data to get times and locations for a set of pictures.

To use with iPhoto, you'll need to SelectAll > File > Export all the pictures somewhere.

Since there's probably a lot of images without EXIF data, all EXIF errors are ignored.

Output is a CSV of the form |"name"|"time"|"lat"|"long"|.
The CSV __doesn't__ have a header row. It formats `time` as golang `time.UnixTime`. 


:cat2: :footprints:
::paw:a
