BYTEMAN_HOME=$BYTEMAN_HOME
BYTEMAN_RULE=$(pwd)/rule.btm
BYTEMAN_OPTS="-javaagent:$BYTEMAN_HOME/lib/byteman.jar=listener:true,boot:$BYTEMAN_HOME/lib/byteman.jar"
if [ "x$BYTEMAN_RULE" != "x" ]; then
   BYTEMAN_OPTS="${BYTEMAN_OPTS},script:$BYTEMAN_RULE"
fi
BYTEMAN_OPTS="$BYTEMAN_OPTS -Dorg.jboss.byteman.transform.all -Dorg.jboss.byteman.debug"

JAVA_OPTS="$BYTEMAN_OPTS $JAVA_OPTS"

java $JAVA_OPTS nak3.com.Sample
