#! /usr/bin/awk -f
{if($1>prev&&NR>1){increased++;}prev=$1;if(NR>3){window=minus2+minus1+$1;if(window>lastWindow){increased2++;}lastWindow=window;}minus2=minus1;minus1=$1}END{print"Part 1: "increased"\nPart 2: "increased2}