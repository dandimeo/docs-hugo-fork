#!/bin/sh
set -e

if [ -z "$ARANGO_INIT_PORT" ] ; then
    ARANGO_INIT_PORT=8529
fi

AUTHENTICATION="true"

# if command starts with an option, prepend arangod
case "$1" in
    -*) set -- arangod "$@" ;;
    *) ;;
esac

# check for numa
NUMACTL=""

if [ -d /sys/devices/system/node/node1 -a -f /proc/self/numa_maps ]; then
    if [ "$NUMA" = "" ]; then
        NUMACTL="numactl --interleave=all"
    elif [ "$NUMA" != "disable" ]; then
        NUMACTL="numactl --interleave=$NUMA"
    fi

    if [ "$NUMACTL" != "" ]; then
        if $NUMACTL echo > /dev/null 2>&1; then
            echo "using NUMA $NUMACTL"
        else
            echo "cannot start with NUMA $NUMACTL: please ensure that docker is running with --cap-add SYS_NICE"
            NUMACTL=""
        fi
    fi
fi

if [ "$1" = 'arangod' ]; then
    # /var/lib/arangodb3 and /var/lib/arangodb3-apps must exist and
    # be writable by the user under which we run the container.

    # Make a copy of the configuration file to patch it, note that this
    # must work regardless under which user we run:
    cp /etc/arangodb3/arangod.conf /tmp/arangod.conf

    ARANGO_STORAGE_ENGINE=rocksdb
    if [ ! -z "$ARANGO_ENCRYPTION_KEYFILE" ]; then
        echo "Using encrypted database"
        sed -i /tmp/arangod.conf -e "s;^.*encryption-keyfile.*;encryption-keyfile=$ARANGO_ENCRYPTION_KEYFILE;"
    fi

    if [ ! -f /var/lib/arangodb3/SERVER ] && [ "$SKIP_DATABASE_INIT" != "1" ]; then
        if [ ! -z "$ARANGO_ROOT_PASSWORD_FILE" ]; then
            if [ -f "$ARANGO_ROOT_PASSWORD_FILE" ]; then
                ARANGO_ROOT_PASSWORD="$(cat $ARANGO_ROOT_PASSWORD_FILE)"
            else
                echo "WARNING: password file '$ARANGO_ROOT_PASSWORD_FILE' does not exist"
            fi
	fi
        # Please note that the +x in the following line is for the case
        # that ARANGO_ROOT_PASSWORD is set but to an empty value, please
        # do not remove!
        if [ -z "${ARANGO_ROOT_PASSWORD+x}" ] && [ -z "$ARANGO_NO_AUTH" ] && [ -z "$ARANGO_RANDOM_ROOT_PASSWORD" ]; then
            echo >&2 'error: database is uninitialized and password option is not specified '
            echo >&2 "  You need to specify one of ARANGO_ROOT_PASSWORD, ARANGO_ROOT_PASSWORD_FILE, ARANGO_NO_AUTH and ARANGO_RANDOM_ROOT_PASSWORD"
            exit 1
        fi

        if [ ! -z "$ARANGO_RANDOM_ROOT_PASSWORD" ]; then
            ARANGO_ROOT_PASSWORD=$(pwgen -s -1 16)
            echo "==========================================="
            echo "GENERATED ROOT PASSWORD: $ARANGO_ROOT_PASSWORD"
            echo "==========================================="
        fi

        if [ ! -z "${ARANGO_ROOT_PASSWORD+x}" ]; then
            echo "Initializing root user...Hang on..."
            ARANGODB_DEFAULT_ROOT_PASSWORD="$ARANGO_ROOT_PASSWORD" /usr/sbin/arango-init-database -c /tmp/arangod.conf --server.rest-server false --log.level error --database.init-database true || true
            export ARANGO_ROOT_PASSWORD

            if [ ! -z "${ARANGO_ROOT_PASSWORD}" ]; then
                ARANGOSH_ARGS=" --server.password ${ARANGO_ROOT_PASSWORD} "
            fi
        else
            ARANGOSH_ARGS=" --server.authentication false"
        fi
    fi

    # if we really want to start arangod and not bash or any other thing
    # prepend --authentication as the FIRST argument
    # (so it is overridable via command line as well)
    shift

    if [ ! -z "$ARANGO_NO_AUTH" ]; then
	    AUTHENTICATION="false"
    fi

    set -- arangod "$@" --server.authentication="$AUTHENTICATION" --config /tmp/arangod.conf
else
    NUMACTL=""
fi

exec $NUMACTL "$@"
