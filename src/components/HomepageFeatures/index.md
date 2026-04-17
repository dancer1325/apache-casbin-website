# Hybrid access control models

* CONF files
  * uses
    * define -- , based on the PERM metamodel (Policy, Effect, Request, Matchers), -- access control models 
  * if you want to change OR upgrade your authorization mechanism -> modify the configuration file

# Flexible policy storage

* Apache Casbin policies
  * 👀are stored👀 
    * ways 
      * memory
      * files
      * databases
    * -- through -- [storage backends OR adapters](/docs/Adapters.md)
      * separated -- from the -- MAIN library
        * ⚠️EXCEPT FOR: default file adapter⚠️
        * Reason:🧠keep the library lightweight🧠
      * supported
        * MySQL
        * Postgres
        * Oracle MongoDB
        * Redis
        * Cassandra
        * AWS S3
      * SOME support filtered policy loading
        * == Apache Casbin load ONLY -- , based on specified filters, -- a subset of policies | storage
        * use cases
          * applications / large-scale & multi-tenant
            * Reason:🧠if you load ALL policies DIRECTLY -> inefficient🧠

# Cross-languages & cross-platforms

* [here](../LanguageIntegration/index.md)
