# $ awk -f preprocessing.awk myfile
BEGIN {
	print "preprocessing"
	
	}
	{print $1="pre"; print $0}
