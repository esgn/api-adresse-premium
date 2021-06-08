#!/bin/bash

export PGUSER=$POSTGRES_USER
export PGDATABASE=$POSTGRES_DB
export PGPASSWORD=$POSTGRES_PASSWORD

psql -c "DROP SCHEMA IF EXISTS $POSTGRES_SCHEMA CASCADE;"
psql -c "CREATE SCHEMA $POSTGRES_SCHEMA;"
psql -c "CREATE EXTENSION IF NOT EXISTS postgis;"

cd /tmp/adresse-premium/
append=''
src_epsg=2154
dst_epsg=4326
for f in *.7z; do
    echo $f
    if [[ $f =~ "RGAF09UTM20" ]]
    then
        src_epsg=5490
    elif [[ $f =~ "RGM04UTM38S" ]]
    then
        src_epsg=4471
    elif [[ $f =~ "RGR92UTM40S" ]]
    then
        src_epsg=2975
    elif [[ $f =~ "UTM22RGFG95" ]]
    then
        src_epsg=2972
    fi
    7z x $f
    xdir=`basename $f .7z`
    cd `find $xdir -type d -name "A_ADR-PARC"`
    shp=`find . -type f -name *.shp`
    shp2pgsql -s $src_epsg:$dst_epsg -D $append $shp $POSTGRES_SCHEMA.adr_parc | psql
    cd /tmp/adresse-premium/
    cd `find $xdir -type d -name "B_ADR-BATI"`
    shp=`find . -type f -name *.shp`
    shp2pgsql -s $src_epsg:$dst_epsg -D $append $shp $POSTGRES_SCHEMA.adr_bati | psql
    cd /tmp/adresse-premium/
    cd `find $xdir -type d -name "C_BATI-PARC"`
    shp=`find . -type f -name *.shp`
    shp2pgsql -s $src_epsg:$dst_epsg -D $append $shp $POSTGRES_SCHEMA.bati_parc | psql
    cd /tmp/adresse-premium/
    rm -rf $xdir
    #rm $f
    append='-a'
done

psql -c "CREATE INDEX adr_parc_id_idx ON $POSTGRES_SCHEMA.adr_parc (id)"
psql -c "CREATE INDEX adr_parc_id_adr_idx ON $POSTGRES_SCHEMA.adr_parc (id_adr)"
psql -c "CREATE INDEX adr_bati_id_idx ON $POSTGRES_SCHEMA.adr_bati (id)"
psql -c "CREATE INDEX adr_bati_id_adr_idx ON $POSTGRES_SCHEMA.adr_bati (id_adr)"
psql -c "CREATE INDEX bati_parc_id_idx ON $POSTGRES_SCHEMA.bati_parc (id)"
psql -c "CREATE INDEX bati_parc_id_bat_idx ON $POSTGRES_SCHEMA.bati_parc (id_bat)"
psql -c "CREATE INDEX bati_parc_idu_idx ON $POSTGRES_SCHEMA.bati_parc (idu)"
