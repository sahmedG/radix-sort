range=$1
case=$2
declare -i counter=$range
numbers=""
numbersArr=()

for ((i = 1; i <= $range; i++)); do
  declare -i numberToAppend="$((RANDOM % 1000))"
  repeats=no
  while [[ "$repeats" = "no" ]]; do
    repeats=yes
    for y in "${numbersArr[@]}"; do
      if [ "$y" = "$numberToAppend" ]; then
        repeats=no
        numberToAppend=$(($y + 1))
      fi
    done
  done
  numbers+="$numberToAppend"
  numbersArr+=("$numberToAppend")
  if [ "$i" != "$counter" ]; then
    numbers+=" "
  fi
done

if [ "$case" = 0 ]; then
  ./push-swap "$numbers"
else
  echo "$numbers"
  ARG=('4 67 3 87 23'); ./push-swap "$ARG" | ./checker "$ARG"
fi
