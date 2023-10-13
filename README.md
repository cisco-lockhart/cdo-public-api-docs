# Landing Page

The landing page is made with XML widgets and highlights the developer resources (docs, learning labs, sandbox, sample code, blogs, etc.) in a  graphical web page.
To edit the text and links, edit /en/index.xml. 
To remove or add a widget, edit index.html
For complete instructions, check out the [User Guide](https://pubhub.cisco.com/project/preview/587/new/)



## Local preview tool

To preview changes locally, follow these instructions to get setup: https://wwwin-github.cisco.com/DevNet/Pubhub-Site-Engine#local-preview-tool

Then, run the following command:

```
curl -O https://wwwin-github.cisco.com/raw/DevNet/Pubhub-Site-Engine/update-local-preview/docker-compose.yml && BUILD_TYPE=doc docker-compose up
```

You can then preview the landing page here: [https://0.0.0.0:9999/](https://0.0.0.0:9999/)