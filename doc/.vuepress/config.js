module.exports = ctx => ({
    dest: 'dist',
    locales: {
        '/': {
            lang: 'zh-CN',
            title: 'Gex',
            description: 'Golang外部程序调用扩展库'
        }
    },
    head: [
        ['link', {rel: 'icon', href: `/logo.png`}],
        ['link', {rel: 'manifest', href: '/manifest.json'}],
        ['link', {rel: "icon", type: "image/png", sizes: "32x32", href: "/icons/favicon-32x32.png"}],
        ['link', {rel: "icon", type: "image/png", sizes: "16x16", href: "/icons/favicon-16x16.png"}],
        ['meta', {name: 'theme-color', content: '#3eaf7c'}],
        ['meta', {name: 'apple-mobile-web-app-capable', content: 'yes'}],
        ['meta', {name: 'apple-mobile-web-app-status-bar-style', content: 'black'}],
        ['link', {rel: 'apple-touch-icon', href: `/icons/favicon-152x152.png`}],
        ['link', {rel: 'mask-icon', href: '/icons/safari-pinned-tab.svg', color: '#3eaf7c'}],
        ['meta', {name: 'msapplication-TileImage', content: '/icons/favicon-144x144.png'}],
        ['meta', {name: 'msapplication-TileColor', content: '#000000'}]
    ],
    markdown: {
        lineNumbers: true
    },
    theme: 'reco',
    themeConfig: {
        logo: '/icons/logo.png',
        repo: 'goexl/gex',
        docsDir: 'doc',
        editLinks: true,
        record: '蜀ICP备2021013439号-3',
        recordLink: 'https://beian.miit.gov.cn/',
        // cyberSecurityRecord: '公安部备案文案',
        // cyberSecurityLink: '公安部备案指向链接',
        // 项目开始时间，只填写年份
        startYear: '2021',
        algolia: ctx.isProd ? ({
            apiKey: 'f7edee00640ed06f44542bb62b4d9e5b',
            indexName: 'doc',
            algoliaOptions: {
                facetFilters: ['tags:v1']
            }
        }) : null,
        smoothScroll: true,
        locales: {
            '/': {
                label: '简体中文',
                selectText: '选择语言',
                ariaLabel: '选择语言',
                editLinkText: '在GitHub上编辑此页',
                lastUpdated: '上次更新',
                nav: require('./nav/zh'),
                sidebar: {
                    '/api/': getApiSidebar(),
                    '/guide/': getGuideSidebar('指南', '深入'),
                    '/checker/': getCheckerSidebar('基础', '进阶'),
                    '/collector/': getCollectorSidebar('基础', '进阶'),
                    '/notifier/': getNotifierSidebar('基础', '进阶'),
                    '/pipe/': getPipeSidebar('使用', '配置'),
                }
            }
        }
    },
    plugins: [
        ['sitemap', {
            hostname: 'https://gex.storezhang.tech',
            exclude: ['/404.html']
        }],
        ['@vuepress/last-updated', {
            transformer: (timestamp, lang) => {
                const moment = require('moment')
                moment.locale(lang)
                const datetime = moment(timestamp)

                return datetime.format('yyyy-MM-DD hh:mm:ss')
            }
        }],
        ['@vuepress/back-to-top', true],
        ['@vuepress/pwa', {
            serviceWorker: true,
            updatePopup: true
        }],
        ['@vuepress/medium-zoom', true],
        ['@vuepress/google-analytics', {
            ga: 'G-RJRHHSRXT8'
        }],
        ['container', {
            type: 'vue',
            before: '<pre class="vue-container"><code>',
            after: '</code></pre>'
        }],
        ['container', {
            type: 'upgrade',
            before: info => `<UpgradePath title="${info}">`,
            after: '</UpgradePath>'
        }],
        ['flowchart']
    ],
    extraWatchFiles: [
        '.vuepress/nav/en.js',
        '.vuepress/nav/zh.js'
    ]
})

function getApiSidebar() {
    return [
        'cli',
        'node'
    ]
}

function getGuideSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'start',
            'args',
            'dir',
            'envs',
            'async',
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'context',
            'quiet',
            'pwe',
            'terminal',
            'stdout',
            'stderr',
        ]
    }]
}

function getCheckerSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'contains',
            'equal',
            'contains.all',
            'contains.any',
            'path.match',
            'regexp',
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'concept',
            'interface',
            'usage',
        ]
    }]
}

function getCollectorSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'options',
            'string',
            'file',
            'writer',
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'concept',
            'interface',
            'usage',
        ]
    }]
}

function getNotifierSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'func',
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'concept',
            'interface',
            'usage',
        ]
    }]
}

function getPipeSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'concept',
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'args',
            'dir',
            'envs',
        ]
    }]
}
