killall polybar

PRIMARY=$(xrandr --query | grep " connected" | grep "primary" | cut -d" " -f1)
OTHERS=$(xrandr --query | grep " connected" | grep -v "primary" | cut -d" " -f1)

# Launch on primary monitor
MONITOR=$PRIMARY polybar --reload example &
MONITOR=$PRIMARY polybar --reload center &
MONITOR=$PRIMARY polybar --reload right &
sleep 1

# Launch on all other monitors
for m in $OTHERS; do
 MONITOR=$m polybar --reload example &
 MONITOR=$m polybar --reload center &
 MONITOR=$m polybar --reload right-np &
done

