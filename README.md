# api-adresse-premium

API REST minimale pour diffusion simple du produit IGN Adresse Premium : https://geoservices.ign.fr/ressources_documentaires/Espace_documentaire/BASES_VECTORIELLES/ADRESSE_PREMIUM/DC_ADRESSE_PREMIUM_3-0.pdf

## Préalables

1. Changer le mot de passe par défaut de la base (`password`) dans le fichier docker-compose.yml
2. Adapter les options de configuration de PostgreSQL dans le fichier docker-compose.yml si nécessaire

## Déploiement

1. Téléchargement du produit Adresse Premium

    `python3 download-dataset.py`

2. Lancement des containers

    `docker-compose up -d`

3. Import des données en base (opération longue pouvant être lancée dans un `screen` par exemple)

    `docker exec -ti adr-premium-postgis /bin/bash /tmp/scripts/import-data.sh`

## Utilisation

Les résultats sont fournis au format GeoJSON.

* (GET) `/adr_parc/{id}` : Récupération d'un lien adresse-parcelle par identifiant (`ADR_PARC[0-9]+`)
* (GET) `/adr_parc/findByAdrId?id=` : Recherche de liens adresse-parcelle par identifiant d'adresse IGN (`ADRNIVX_[0-9]+`)
* (GET) `/adr_bati/{id}` : Récupération d'un lien adresse-batiment par identifiant (`ADR_BATI[0-9]+`)
* (GET) `/adr_bati/findByAdrId?id=` : Recherche de liens adresse-batiment par identifiant d'adresse IGN (`ADRNIVX_[0-9]+`)
* (GET) `/bati-parc/{id}` : Récupération d'un lien batiment-parcelle par identifiant (`BAT_PARC[0-9]+`)
* (GET) `/bati-parc/findByBatId?id=` : Recherche de liens batiment-parcelle par identifiant de batiment IGN (`BATIMENT[0-9]+`)
* (GET) `/bati-parc/findByParcId?id=` : Recherche de liens batiment-parcelle par identifiant de parcelle

## Arrêt du service

* Sans suppression des données en base : `docker-compose down`
* Avec suppression des données en base : `docker-compose down -v`
