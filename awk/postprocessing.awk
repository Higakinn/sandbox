# awk -f postprocessing.awk myfile
{print $1="END"; print $0}
END { print "END!!!!!"}
