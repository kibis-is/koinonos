/**
 * @type {import('semantic-release').GlobalConfig}
 */
export default {
  branches: [
    'main',
    {
      name: 'beta',
      prerelease: true,
    },
  ],
  plugins: [
    '@semantic-release/commit-analyzer',
    '@semantic-release/release-notes-generator',
    '@semantic-release/changelog',
    [
      '@semantic-release/exec',
      {
        prepareCmd:
          './scripts/update_version.sh ${nextRelease.version} && ./scripts/update_issue_templates.sh ${nextRelease.version}',
        publishCmd: [
          // os x
          'darwin-amd64',
          'darwin-arm64',
          // linux
          'linux-amd64',
          'linux-arm64',
          // windows
          'windows-amd64',
          'windows-arm64',
        ].map((platform) => `./scripts/build.sh "${platform}"`).join(' && '),
      },
    ],
    [
      '@semantic-release/git',
      {
        assets: [
          '.github/ISSUE_TEMPLATE/bug_report_template.yml',
          'CHANGELOG.md',
          'VERSION',
        ],
        message: 'chore(release): ${nextRelease.version}\n\n${nextRelease.notes}',
      },
    ],
    [
      '@semantic-release/github',
      {
        assets: ['dist/*.tar.gz'],
        releasedLabels: ['🚀 released'],
      },
    ],
  ],
};
