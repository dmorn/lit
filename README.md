# lit
Literature review tool. Supports Elsevier's Scopus.

# Usage
Three tools are provided to help researchers perform the first phases of a
sistematic literature review (SLR). Everything starts with `lit-max`: its the
query refinement phase in which the query is exposed to the library and the
number of hits is returned. It is suggested to tune the query till ~500 results
are returned. The second phase is the download one, where the results are
acually downloaded and stored locally. This is performed using the `lit-get`
tool. The third phase is the time-demanding one, in which each publication is
reviewed and the researcher is supposed to accept/reject papers based on some
exclusion/inclusion criteria. This is done thourgh `lit-review`.

# Features
The `lit-*` suite uses an event-based database (single file selected through
the -edb flag) to store everything. Just ensure you don't loose this file and
you'll be fine. For now, you can freely edit the file. In the future, I plan on
using a markov-chain strategy to ensure other researchers that your results
have not been tampered, but this is going to happen only if the tool exits the
"toy/prototype" phase and someone else is using it.

For feature requests and everything else, open an issue.

# Recovering
The program's state is constructed from its .edb file, by default lit.edb. If
something goes wrong, users are invited to open it up and edit its contents for
now, for example by deleting one or more reviews.
