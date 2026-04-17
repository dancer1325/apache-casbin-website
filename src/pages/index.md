* [documentation](../../docs)
* [supported languages](../components/LanguageIntegration/index.md)
* [features](../components/HomepageFeatures/index.md)
* [Apache Casbin Online Editor](../components/EditorPreview/index.md)
* 
* TODO: 


function RoleManager() {
  return (
    <div className={styles.rolemanager}>
      <div className="container text--center">
        <div className="row">
          <div className="col" style={{marginBlock: "auto"}}>
            <h3><Translate>Role Manager</Translate></h3>
            <Translate
              values={{
                rolemanagersLink: (
                  <Link to="/docs/role-managers">
                    <Translate>role-managers</Translate>
                  </Link>
                ),
              }}
            >
              {"The role manager handles RBAC role hierarchy (user-role mappings) in Apache Casbin. It can load role data from Apache Casbin policy rules or from external sources like LDAP, Okta, Auth0, Azure AD, etc. To keep the library lightweight, role manager code is separated from the main library (except for the default one). See all available {rolemanagersLink}."}
            </Translate>
          </div>
          <div style={{marginInline: "auto"}}>
            <img src="/img/role.png" alt="Role manager" width="500" height="500" />
          </div>
        </div>
      </div>
    </div>
  );
}

function Showcase() {
  return (
    <div className="hero text--center showcase">
      <div className="container">
        <h1><Translate>Who&apos;s using Apache Casbin?</Translate></h1>
        <p style={{
          width: "50vw",
          margin: "auto",
        }}>
          <Translate values={{
            UsersPage: (
              <Link to="/users">
                <Translate>check out these apps</Translate>
              </Link>
            ),
          }}>
            {"Hundreds of projects use Apache Casbin, from Fortune 500 companies to new startups. If you want to see what can be built with Apache Casbin, {UsersPage}!"}
          </Translate>
        </p>
        <br /><br />
        <UserList />
      </div>
    </div>
  );
}

export default function Home() {
  return (
    <Layout
      title="Apache Casbin · An authorization library"
      description="An authorization library that supports access control models like ACL, RBAC, ABAC, ReBAC, PBAC, OrBAC, BLP, Biba, LBAC, UCON, Priority, RESTful for Golang, Java, C/C++, Node.js, Javascript, PHP, Laravel, Python, .NET (C#), Delphi, Rust, Ruby, Swift (Objective-C), Lua (OpenResty), Dart (Flutter) and Elixir">
      <main>
        <RoleManager />
        <Showcase />
      </main>
    </Layout>
  );
}
