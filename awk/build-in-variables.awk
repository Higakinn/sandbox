# awk 'BEGIN{FS=":"; OFS="-"} {print $1,$6,$7}' /etc/passwd
# OFS (output file seperator): default→space, 
# FS (file seperator): 
BEGIN{FS=":"; OFS="-"} {print $1,$6,$7}
