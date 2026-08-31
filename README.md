Voici le README complet traduit en français, prêt à copier-coller :
Markdown# Crédit à Paranørmal (BreakArchive) et Dopa
# System Monitor Discord Rich Presence

Un script Python moderne pour surveiller et afficher l'utilisation du CPU, de la mémoire, des E/S disque et des E/S réseau via Discord Rich Presence.

## Fonctionnalités

- Affiche les statistiques système en direct (CPU, RAM, Disque, Réseau) dans votre statut Discord
- Alterne entre différentes combinaisons de statistiques pour une présence dynamique
- Formatage lisible de toutes les tailles (ex. : Mo, Go)
- Simple, sans logs, sans fichiers de sortie, uniquement Discord RPC

## Prérequis

- Python 3.7+
- [psutil](https://pypi.org/project/psutil/)
- [pypresence](https://pypi.org/project/pypresence/)

## Installation

```sh
pip install -r requirements.txt
Utilisation
shpython discord.py
Arguments optionnels

--rpc-update-interval : Fréquence de mise à jour du statut Discord (en secondes, par défaut : 10)
--discord-client-id : Votre Client ID d'application Discord (par défaut : fourni dans le script)

Exemple :
shpython discord.py --rpc-update-interval 5 --discord-client-id 123456789012345678
Notes

Discord doit être lancé sur votre ordinateur pour que le Rich Presence fonctionne.
Si vous souhaitez utiliser vos propres icônes/images, configurez-les dans le Discord Developer Portal et utilisez votre propre Client ID.

Licence
MIT License# Crédit to Paranørmal (BreakArchive) and Dopa
# System Monitor Discord Rich Presence

A modern Python script to monitor and display CPU, memory, disk I/O, and network I/O via Discord Rich Presence.

## Features

- Shows live system stats (CPU, RAM, Disk, Network) in your Discord status
- Rotates between different stat combinations for a dynamic presence
- Human-readable formatting for all sizes (e.g., MB, GB)
- Simple, no-logging, no-file-output, just Discord RPC

## Requirements

- Python 3.7+
- [psutil](https://pypi.org/project/psutil/)
- [pypresence](https://pypi.org/project/pypresence/)

## Installation

```sh
pip install -r requirements.txt
```

## Usage

```sh
python discord.py
```

### Optional arguments

- `--rpc-update-interval`: How often to update Discord status (seconds, default: 10)
- `--discord-client-id`: Your Discord Application Client ID (default: provided in script)

Example:

```sh
python discord.py --rpc-update-interval 5 --discord-client-id 123456789012345678
```

## Notes

- You must have Discord running on your computer for Rich Presence to work.
- If you want your own icon/images, set them up in your Discord Developer Portal and use your own Client ID.

## License

MIT License
