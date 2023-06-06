### Get stuff installed

To work on the docs you need to first get the cli tool to build docs:

- Make sure you have nvm installed and that you use node 16
- Clone the GetStream/stream-chat-docusaurus-cli repo
- Switch to the staging branch
- Make sure to run `nvm use 16`
- Run `yarn install`
- Run `npm install -g` to get `stream-docusaurus-cli` in npm binaries cache

### Running docs server

- make sure to cd the `docusaurus/video` path of this repo
- `nvm use 16` otherwise you get stupid node errors
- `npx stream-chat-docusaurus -s` starts a developer server 

### Basic stuff

- documentation sections are paths under docs/api
- documentation is inside mdx files, that's a fancy extension of Markdown
- mdx files must contain the intro section (the --- bit) otherwise they will not get picked up by the server
- for code examples you can use `Tabs` and `TabItem`, see here for [more](https://docusaurus.io/docs/markdown-features/code-blocks#multi-language-support-code-blocks) 

### Other Notes
- code examples with tabs are super dumb, you need to take into account that MDX is crap (indentation new lines will break the parser and lead to errors)
- the development server does not support all sorts of reload, really often if a page breaks you need to ctrl-c and restart it 
