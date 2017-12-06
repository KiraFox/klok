Klok
====

Klok is a simple command line time tracker.  Time is tracked by clocking in and
out. You can check your times by the current day or week.


Installing
----------

To install this tool you need to have [Go][go] and [Git][git] installed.  You
also need to have `$GOPATH/bin` included in your `PATH` environment variable.

Run:

	$ go get -u github.com/KiraFox/klok

**Note:** _This same command can be used to update your copy of Klok._

[go]: https://golang.org/
[git]: https://git-scm.com/


Usage
-----

Klok is used in the command line and is controlled by using subcommands.

The following is a list of the available subcommands:
	
- `klok in` : Clocks you in with the current time.
- `klok out` : Clocks you out with the current time.
- `klok today` : Shows you the total time you spent clocked in today.
- `klok week` : Shows you the summary of clocked in time for the week.
- `klok edit` : Opens the file where your time is stored for the current week.

The editor run by `klok edit` can be set with your `EDITOR` environment
variable.  If a preference is not set, Klok will try to use your default editor.


Storage
-------

Times are stored in different locations depending on your OS.

- windows : `%LOCALAPPDATA%\klok`
- others : `$HOME/.local/share/klok`

Files are named based on the year and week number.  For example, `2017-wk22.txt`
