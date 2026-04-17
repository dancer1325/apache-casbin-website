# [Apache Casbin website](https://casbin.apache.org/)

* goal
  * Apache Casbin's documentation site's source
    * -- based on -- [Docusaurus](https://docusaurus.io/) v3

* Apache Casbin
  * == [Apache Software Foundation (ASF)](https://www.apache.org/)'s [incubation](https://incubator.apache.org/projects/casbin.html) project
  * == ⭐️access control library⭐️ /
    * open-source
    * powerful
    * efficient
    * support
      * 👀MULTIPLE authorization models👀
    * uses |
      * applications
      * AI gateway
      * MCP
      * clouds
      * web apps

## documentation
* [here](src/pages/index.md)
* [blog](blog)

## Build and deploy

TODO: 
On push to `master` or manual dispatch, the [Build and Deploy](https://github.com/apache/casbin-website/actions/workflows/master.yml) workflow:

1. Installs dependencies and, for non–pull-request runs, runs Crowdin sync (`yarn crowdin:sync`; requires the `CROWDIN_PERSONAL_TOKEN` repository secret).
2. Runs `yarn build` to produce the static site.
3. Publishes the `build/` output to the **`asf-site`** branch of this repository (via [peaceiris/actions-gh-pages](https://github.com/peaceiris/actions-gh-pages)).

ASF infrastructure serves the site from `asf-site` as configured in [`.asf.yaml`](./.asf.yaml). **Pull requests** run the workflow without Crowdin sync or publishing to `asf-site`.

Additional checks (Markdown lint, doc link checks, etc.) live under [`.github/workflows/`](./.github/workflows/).

## Local development

### Requirements

1. [Git](https://git-scm.com/downloads)
2. [Node.js](https://nodejs.org/) **20.x or later** (matches `engines` in `package.json` and CI)
3. [Yarn Classic (v1)](https://classic.yarnpkg.com/en/docs/install)

### Run locally

```bash
git clone https://github.com/apache/casbin-website.git
cd casbin-website
yarn install
yarn start
```

Other useful scripts: `yarn build`, `yarn serve`, `yarn lint`.

## Contributing

You can help with documentation, translations, and front-end work. Mailing lists and project metadata are listed on the [incubator status page](https://incubator.apache.org/projects/casbin.html) (for example `dev@casbin.apache.org`). On-site contributing guidance is also available under **Docs → Start Contributing** on the published site.

### Documentation writing

For the sidebar, see [Docusaurus: Sidebar](https://docusaurus.io/docs/sidebar). For Markdown/MDX features, see [Markdown Features](https://docusaurus.io/docs/markdown-features).

A typical doc front matter and structure:

````md
---
title: Title
description: description
keywords: [keyword1, keyword2]
authors: [GitHub username]
---

## Headers

Only h2 and h3 will be in the TOC by default, so h1 is not recommended to use.

### h3

content

#### h4

content

````

We use [markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli) for Markdown/MDX:

```bash
yarn lint:md
```

For VS Code, Sublime, or Vim/Neovim, consider a [markdownlint extension](https://github.com/DavidAnson/markdownlint#related).

#### Caution

##### Admonitions

You can use [Admonitions](https://docusaurus.io/docs/markdown-features/admonitions); leave blank lines around blocks, for example:

```md
:::info Title

Title is optional

:::
```

##### JSX

You can use JSX in docs (e.g. [Tabs](https://docusaurus.io/docs/markdown-features/tabs)). To avoid Crowdin breaking code, follow [mdx-solutions](https://docusaurus.io/docs/i18n/crowdin#mdx-solutions) and wrap JSX with `mdx-code-block`:

````md
```mdx-code-block
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
```

```mdx-code-block
<Tabs>
<TabItem value="go" label="Go">
```

content

```mdx-code-block
</TabItem>
<TabItem value="java" label="Java">
```

content

```mdx-code-block
</TabItem>
</Tabs>
```

````

### Translation

[Crowdin](https://crowdin.com/project/casbin-website) and [Docusaurus i18n](https://docusaurus.io/docs/i18n/introduction) are used for translations.

Do not translate directive strings like `:::note` or `:::tip`; wrong translations can break rendering.

Do not translate interpolation placeholders in sentences such as:

```text
At our {repoLink}, browse and submit {issueLink} or {prLink} for bugs you find or any new features you may want implemented.
```

Keep `{repoLink}`, `{issueLink}`, etc. as-is; see [translate props](https://docusaurus.io/docs/docusaurus-core#translate-props).

Do not translate metadata lines like `authors: [casbin]`.

