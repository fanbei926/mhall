
running()
{
  if [ -f "$1" ]
  then
    local PID=$(cat "$1" 2>/dev/null) || return 1
    kill -0 "$PID" 2>/dev/null
    return
  fi
  rm -f "$1"
  return 1
}


PIDFILE=JAVAPID
if running "$PIDFILE"
	then
		pid=`head -1 $PIDFILE`
		echo Aready Running "$pid"
		exit 1
fi 


JAVA_OPTIONS="${JAVA_OPTIONS} -server -Xms2g -Xmx4g -Xmn2g -Xss256k -XX:PermSize=32m -XX:MaxPermSize=32m"
JAVA_OPTIONS="${JAVA_OPTIONS} -XX:+UseConcMarkSweepGC -XX:+UseParNewGC -XX:+UseCMSCompactAtFullCollection  -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=70"
JAVA_OPTIONS="${JAVA_OPTIONS} -XX:+CMSParallelRemarkEnabled -XX:SoftRefLRUPolicyMSPerMB=0 -XX:+CMSClassUnloadingEnabled -XX:SurvivorRatio=8 -XX:+DisableExplicitGC"
JAVA_OPTIONS="${JAVA_OPTIONS} -verbose:gc -Xloggc:./logs/server_gc.log -XX:+PrintGCDetails -XX:+PrintGCDateStamps"

echo ${JAVA_OPTIONS}
nohup java ${JAVA_OPTIONS} -jar ghall-web-server-0.0.9-SNAPSHOT.jar  >/dev/null  2>&1 &
rm -f "$PIDFILE"
echo $! > "$PIDFILE"
# tail -F logs/app.log