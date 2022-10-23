# echo 'AAABBC' | awk -f build-in-variables2.awk
BEGIN{FIELDWIDTHS="3 2 1"}{print $1,$2,$3}
