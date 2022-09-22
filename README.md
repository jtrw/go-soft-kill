# Soft kill
Simple utility for soft kill processes for `php` scripts.

We will be us `kill -SIGTERM pid` and send it posix signal to `php` script. In php script
we use the library [Posix Signal Handler](https://github.com/jtrw/posix-signal-handler)
which check for the posix signal and terminate our script.
