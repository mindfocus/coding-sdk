# \DefaultApi

All URIs are relative to *https://e.coding.net/open-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveTestRun**](DefaultApi.md#ArchiveTestRun) | **Post** /?action&#x3D;ArchiveTestRun | 测试计划归档
[**AttachResourceScopeToPolicy**](DefaultApi.md#AttachResourceScopeToPolicy) | **Post** /?action&#x3D;AttachResourceScopeToPolicy | 权限组添加可用的资源，原有其它资源不受影响，若已存在的资源不再进行追加
[**AttachToResource**](DefaultApi.md#AttachToResource) | **Post** /?action&#x3D;AttachToResource | 授权追加，原有其它授权不受影响，若授权已存在不再进行追加
[**BindCdApplicationToProject**](DefaultApi.md#BindCdApplicationToProject) | **Post** /?action&#x3D;BindCdApplicationToProject | 绑定 CD 应用到项目
[**BoundExternalDepot**](DefaultApi.md#BoundExternalDepot) | **Post** /?action&#x3D;BoundExternalDepot | 外部仓库关联
[**CancelCdPipeline**](DefaultApi.md#CancelCdPipeline) | **Post** /?action&#x3D;CancelCdPipeline | CD 部署流程取消执行
[**ClearCodingCIJobCache**](DefaultApi.md#ClearCodingCIJobCache) | **Post** /?action&#x3D;ClearCodingCIJobCache | 构建计划缓存清理
[**CreateArtifactCredit**](DefaultApi.md#CreateArtifactCredit) | **Post** /?action&#x3D;CreateArtifactCredit | 制品授信清单创建
[**CreateArtifactProperties**](DefaultApi.md#CreateArtifactProperties) | **Post** /?action&#x3D;CreateArtifactProperties | 制品属性新增（指定版本）
[**CreateArtifactRepository**](DefaultApi.md#CreateArtifactRepository) | **Post** /?action&#x3D;CreateArtifactRepository | 制品仓库创建
[**CreateAttachmentPrepareSignUrl**](DefaultApi.md#CreateAttachmentPrepareSignUrl) | **Post** /?action&#x3D;CreateAttachmentPrepareSignUrl | 附件预上传信息生成
[**CreateBinaryFiles**](DefaultApi.md#CreateBinaryFiles) | **Post** /CreateBinaryFiles | Git文件-Git仓库创建二进制文件
[**CreateBranchProtection**](DefaultApi.md#CreateBranchProtection) | **Post** /CreateBranchProtection | 仓库设置-新增代码保护规则
[**CreateBranchProtectionMember**](DefaultApi.md#CreateBranchProtectionMember) | **Post** /CreateBranchProtectionMember | 仓库设置-新增保护分支规则管理员
[**CreateCaseResult**](DefaultApi.md#CreateCaseResult) | **Post** /?action&#x3D;CreateCaseResult | 测试用例添加测试结果
[**CreateCdCloudAccount**](DefaultApi.md#CreateCdCloudAccount) | **Post** /?action&#x3D;CreateCdCloudAccount | CD 云账号添加
[**CreateCdHostServerGroup**](DefaultApi.md#CreateCdHostServerGroup) | **Post** /?action&#x3D;CreateCdHostServerGroup | CD 主机组添加
[**CreateCdPipeline**](DefaultApi.md#CreateCdPipeline) | **Post** /?action&#x3D;CreateCdPipeline | CD 部署流程创建
[**CreateCdTask**](DefaultApi.md#CreateCdTask) | **Post** /?action&#x3D;CreateCdTask | CD 任务执行
[**CreateCodingCIJob**](DefaultApi.md#CreateCodingCIJob) | **Post** /?action&#x3D;CreateCodingCIJob | 构建计划创建
[**CreateCodingCIJobByTeamTemplate**](DefaultApi.md#CreateCodingCIJobByTeamTemplate) | **Post** /?action&#x3D;CreateCodingCIJobByTeamTemplate | 构建计划-根据团队模版创建
[**CreateCodingProject**](DefaultApi.md#CreateCodingProject) | **Post** /?action&#x3D;CreateCodingProject | 项目创建
[**CreateDepartment**](DefaultApi.md#CreateDepartment) | **Post** /?action&#x3D;CreateDepartment | 部门创建
[**CreateDepotByTemplate**](DefaultApi.md#CreateDepotByTemplate) | **Post** /CreateDepotByTemplate | 仓库信息-根据模板创建仓库
[**CreateDepotFilePushRule**](DefaultApi.md#CreateDepotFilePushRule) | **Post** /CreateDepotFilePushRule | 仓库设置-新增git仓库文件推送规则
[**CreateDepotFilePushRulePrivilege**](DefaultApi.md#CreateDepotFilePushRulePrivilege) | **Post** /CreateDepotFilePushRulePrivilege | 仓库设置-新增git仓库文件推送规则特权者
[**CreateFile**](DefaultApi.md#CreateFile) | **Post** /?action&#x3D;CreateFile | 文件创建
[**CreateFolder**](DefaultApi.md#CreateFolder) | **Post** /?action&#x3D;CreateFolder | 文件夹创建
[**CreateGitBranch**](DefaultApi.md#CreateGitBranch) | **Post** /CreateGitBranch | 仓库分支-用于代码仓库新建分支
[**CreateGitCommit**](DefaultApi.md#CreateGitCommit) | **Post** /CreateGitCommit | Git提交-创建一次提交
[**CreateGitCommitComment**](DefaultApi.md#CreateGitCommitComment) | **Post** /CreateGitCommitComment | Git提交-为某次提交创建一条评论
[**CreateGitCommitNote**](DefaultApi.md#CreateGitCommitNote) | **Post** /CreateGitCommitNote | Git提交-创建提交注释。注意：对于 NotesRef 参数建议默认为空，因为git会使用默认的ref ：refs/notes/commits，如果填了这个参数，会用这个参数指定的ref来保存您的git note，有可能会覆盖您原有的ref。
[**CreateGitDeployKey**](DefaultApi.md#CreateGitDeployKey) | **Post** /CreateGitDeployKey | 仓库设置-新建部署公钥
[**CreateGitDepot**](DefaultApi.md#CreateGitDepot) | **Post** /?action&#x3D;CreateGitDepot | 仓库信息-创建代码仓库
[**CreateGitFiles**](DefaultApi.md#CreateGitFiles) | **Post** /CreateGitFiles | Git文件-创建仓库文件
[**CreateGitMergeReq**](DefaultApi.md#CreateGitMergeReq) | **Post** /CreateGitMergeReq | 合并请求-创建git合并请求
[**CreateGitMergeRequest**](DefaultApi.md#CreateGitMergeRequest) | **Post** /CreateGitMergeRequest | 合并请求-创建Git合并请求mr
[**CreateGitProtectedTagRule**](DefaultApi.md#CreateGitProtectedTagRule) | **Post** /CreateGitProtectedTagRule | 仓库设置-创建标签保护规则
[**CreateGitProtectedTagRules**](DefaultApi.md#CreateGitProtectedTagRules) | **Post** /CreateGitProtectedTagRules | 仓库设置-批量创建标签保护规则
[**CreateGitRelease**](DefaultApi.md#CreateGitRelease) | **Post** /CreateGitRelease | 版本信息-新建git版本
[**CreateGitTag**](DefaultApi.md#CreateGitTag) | **Post** /CreateGitTag | 标签信息-创建标签
[**CreateIssue**](DefaultApi.md#CreateIssue) | **Post** /CreateIssue | 事项创建
[**CreateIssueBlock**](DefaultApi.md#CreateIssueBlock) | **Post** /?action&#x3D;CreateIssueBlock | 前置事项添加
[**CreateIssueComment**](DefaultApi.md#CreateIssueComment) | **Post** /?action&#x3D;CreateIssueComment | 事项评论创建
[**CreateIssueModule**](DefaultApi.md#CreateIssueModule) | **Post** /CreateIssueModule | 事项模块创建
[**CreateIssueWorkHours**](DefaultApi.md#CreateIssueWorkHours) | **Post** /?action&#x3D;CreateIssueWorkHours | 工时登记
[**CreateIteration**](DefaultApi.md#CreateIteration) | **Post** /?action&#x3D;CreateIteration | 迭代创建
[**CreateMemberSshKey**](DefaultApi.md#CreateMemberSshKey) | **Post** /CreateMemberSshKey | 仓库设置-导入团队成员SSH公钥
[**CreateMergeRequestNote**](DefaultApi.md#CreateMergeRequestNote) | **Post** /CreateMergeRequestNote | 合并请求-创建合并请求行评论和改动文件diff行评论
[**CreateMergeRequestReviewer**](DefaultApi.md#CreateMergeRequestReviewer) | **Post** /CreateMergeRequestReviewer | 合并请求-新增合并请求评审者
[**CreatePolicy**](DefaultApi.md#CreatePolicy) | **Post** /?action&#x3D;CreatePolicy | 权限组创建
[**CreateProgram**](DefaultApi.md#CreateProgram) | **Post** /?action&#x3D;CreateProgram | 项目集创建
[**CreateProgramMemberPolicy**](DefaultApi.md#CreateProgramMemberPolicy) | **Post** /?action&#x3D;CreateProgramMemberPolicy | 项目集成员权限组添加
[**CreateProgramProjects**](DefaultApi.md#CreateProgramProjects) | **Post** /?action&#x3D;CreateProgramProjects | 项目集中添加项目
[**CreateProjectAnnouncement**](DefaultApi.md#CreateProjectAnnouncement) | **Post** /?action&#x3D;CreateProjectAnnouncement | 项目公告创建
[**CreateProjectLabel**](DefaultApi.md#CreateProjectLabel) | **Post** /?action&#x3D;CreateProjectLabel | 项目标签创建
[**CreateProjectMemberPrincipal**](DefaultApi.md#CreateProjectMemberPrincipal) | **Post** /?action&#x3D;CreateProjectMemberPrincipal | 项目成员主体新增(包含用户组、部门、成员)
[**CreateProjectWithTemplate**](DefaultApi.md#CreateProjectWithTemplate) | **Post** /?action&#x3D;CreateProjectWithTemplate | 模版项目创建
[**CreateReadOnlyRef**](DefaultApi.md#CreateReadOnlyRef) | **Post** /?action&#x3D;CreateReadOnlyRef | 仓库分支-创建只读分支
[**CreateRelease**](DefaultApi.md#CreateRelease) | **Post** /CreateRelease | 版本创建
[**CreateReport**](DefaultApi.md#CreateReport) | **Post** /?action&#x3D;CreateReport | 测试报告创建
[**CreateRequirementDefectRelation**](DefaultApi.md#CreateRequirementDefectRelation) | **Post** /?action&#x3D;CreateRequirementDefectRelation | 需求关联缺陷
[**CreateSshKey**](DefaultApi.md#CreateSshKey) | **Post** /?action&#x3D;CreateSshKey | 仓库设置-导入用户SSH公钥
[**CreateTestCase**](DefaultApi.md#CreateTestCase) | **Post** /?action&#x3D;CreateTestCase | 测试用例创建
[**CreateTestCaseSection**](DefaultApi.md#CreateTestCaseSection) | **Post** /?action&#x3D;CreateTestCaseSection | 测试用例分组创建
[**CreateTestDefect**](DefaultApi.md#CreateTestDefect) | **Post** /?action&#x3D;CreateTestDefect | 测试任务关联缺陷
[**CreateTestResult**](DefaultApi.md#CreateTestResult) | **Post** /?action&#x3D;CreateTestResult | 测试任务添加测试结果
[**CreateTestResults**](DefaultApi.md#CreateTestResults) | **Post** /?action&#x3D;CreateTestResults | 测试任务状态批量更新
[**CreateTestRun**](DefaultApi.md#CreateTestRun) | **Post** /?action&#x3D;CreateTestRun | 测试计划创建
[**CreateTestStepResult**](DefaultApi.md#CreateTestStepResult) | **Post** /?action&#x3D;CreateTestStepResult | 测试任务添加某步骤的测试结果
[**CreateUserGroup**](DefaultApi.md#CreateUserGroup) | **Post** /?action&#x3D;CreateUserGroup | 用户组创建
[**CreateUserGroupUsers**](DefaultApi.md#CreateUserGroupUsers) | **Post** /?action&#x3D;CreateUserGroupUsers | 用户组添加用户
[**DeleteAllUsersOnGroup**](DefaultApi.md#DeleteAllUsersOnGroup) | **Post** /?action&#x3D;DeleteAllUsersOnGroup | 用户组清理用户
[**DeleteArtifactProperties**](DefaultApi.md#DeleteArtifactProperties) | **Post** /?action&#x3D;DeleteArtifactProperties | 制品属性删除
[**DeleteBranchProtection**](DefaultApi.md#DeleteBranchProtection) | **Post** /DeleteBranchProtection | 仓库设置-删除保护分支规则
[**DeleteBranchProtectionMember**](DefaultApi.md#DeleteBranchProtectionMember) | **Post** /DeleteBranchProtectionMember | 仓库设置-删除保护分支规则管理员
[**DeleteCdCloudAccount**](DefaultApi.md#DeleteCdCloudAccount) | **Post** /?action&#x3D;DeleteCdCloudAccount | CD 云账号删除
[**DeleteCdHostServerGroup**](DefaultApi.md#DeleteCdHostServerGroup) | **Post** /?action&#x3D;DeleteCdHostServerGroup | CD 主机组删除
[**DeleteCdPipeline**](DefaultApi.md#DeleteCdPipeline) | **Post** /?action&#x3D;DeleteCdPipeline | CD 部署流程删除
[**DeleteCodingCIBuild**](DefaultApi.md#DeleteCodingCIBuild) | **Post** /?action&#x3D;DeleteCodingCIBuild | 构建删除
[**DeleteCodingCIJob**](DefaultApi.md#DeleteCodingCIJob) | **Post** /?action&#x3D;DeleteCodingCIJob | 构建计划删除
[**DeleteDepartment**](DefaultApi.md#DeleteDepartment) | **Post** /?action&#x3D;DeleteDepartment | 部门删除
[**DeleteDepotFilePushRule**](DefaultApi.md#DeleteDepotFilePushRule) | **Post** /DeleteDepotFilePushRule | 仓库设置-删除git仓库文件推送规则
[**DeleteDepotFilePushRuleDenyPrivilege**](DefaultApi.md#DeleteDepotFilePushRuleDenyPrivilege) | **Post** /DeleteDepotFilePushRuleDenyPrivilege | 仓库设置-删除git仓库特权者文件推送权限
[**DeleteGitBranch**](DefaultApi.md#DeleteGitBranch) | **Post** /DeleteGitBranch | 仓库分支-删除代码仓库分支
[**DeleteGitDeployKey**](DefaultApi.md#DeleteGitDeployKey) | **Post** /DeleteGitDeployKey | 仓库设置-删除部署公钥
[**DeleteGitDepot**](DefaultApi.md#DeleteGitDepot) | **Post** /DeleteGitDepot | 仓库信息-删除git仓库
[**DeleteGitFiles**](DefaultApi.md#DeleteGitFiles) | **Post** /DeleteGitFiles | Git文件-删除文件并提交
[**DeleteGitMergedBranches**](DefaultApi.md#DeleteGitMergedBranches) | **Post** /DeleteGitMergedBranches | 仓库分支-删除已合并到默认分支的分支（此操作不会删除受保护的分支)
[**DeleteGitProtectedTagRule**](DefaultApi.md#DeleteGitProtectedTagRule) | **Post** /DeleteGitProtectedTagRule | 标签信息-删除标签保护规则
[**DeleteGitRelease**](DefaultApi.md#DeleteGitRelease) | **Post** /DeleteGitRelease | 版本信息-删除仓库版本
[**DeleteGitTag**](DefaultApi.md#DeleteGitTag) | **Post** /DeleteGitTag | 标签信息-代码仓库删除tag
[**DeleteIssue**](DefaultApi.md#DeleteIssue) | **Post** /?action&#x3D;DeleteIssue | 事项删除
[**DeleteIssueBlock**](DefaultApi.md#DeleteIssueBlock) | **Post** /?action&#x3D;DeleteIssueBlock | 前置事项删除
[**DeleteIssueModule**](DefaultApi.md#DeleteIssueModule) | **Post** /DeleteIssueModule | 事项模块删除
[**DeleteIssueWorkHours**](DefaultApi.md#DeleteIssueWorkHours) | **Post** /?action&#x3D;DeleteIssueWorkHours | 工时日志删除
[**DeleteIteration**](DefaultApi.md#DeleteIteration) | **Post** /?action&#x3D;DeleteIteration | 迭代删除
[**DeleteMemberSshKey**](DefaultApi.md#DeleteMemberSshKey) | **Post** /DeleteMemberSshKey | 仓库设置-删除团队成员的SSH公钥
[**DeleteMergeRequestNote**](DefaultApi.md#DeleteMergeRequestNote) | **Post** /DeleteMergeRequestNote | 合并请求-删除合并请求行评论和改动文件diff行评论
[**DeleteMergeRequestReviewer**](DefaultApi.md#DeleteMergeRequestReviewer) | **Post** /DeleteMergeRequestReviewer | 合并请求-删除mr评审者
[**DeleteOneProject**](DefaultApi.md#DeleteOneProject) | **Post** /?action&#x3D;DeleteOneProject | 单个项目删除
[**DeletePoliciesById**](DefaultApi.md#DeletePoliciesById) | **Post** /?action&#x3D;DeletePoliciesById | 权限组批量删除
[**DeleteProgramMemberPolicy**](DefaultApi.md#DeleteProgramMemberPolicy) | **Post** /?action&#x3D;DeleteProgramMemberPolicy | 项目集成员权限组删除
[**DeleteProjectAnnouncement**](DefaultApi.md#DeleteProjectAnnouncement) | **Post** /?action&#x3D;DeleteProjectAnnouncement | 项目公告删除
[**DeleteProjectLabel**](DefaultApi.md#DeleteProjectLabel) | **Post** /?action&#x3D;DeleteProjectLabel | 项目标签删除
[**DeleteProjectMemberPrincipal**](DefaultApi.md#DeleteProjectMemberPrincipal) | **Post** /?action&#x3D;DeleteProjectMemberPrincipal | 项目成员主体删除(包含用户组、部门、成员)
[**DeleteRelease**](DefaultApi.md#DeleteRelease) | **Post** /DeleteRelease | 版本删除
[**DeleteReport**](DefaultApi.md#DeleteReport) | **Post** /?action&#x3D;DeleteReport | 测试报告删除
[**DeleteRequirementDefectRelation**](DefaultApi.md#DeleteRequirementDefectRelation) | **Post** /?action&#x3D;DeleteRequirementDefectRelation | 需求取消关联缺陷
[**DeleteSshKey**](DefaultApi.md#DeleteSshKey) | **Post** /?action&#x3D;DeleteSshKey | 仓库设置-删除当前用户的SSH公钥
[**DeleteTeamLevelDepotSpec**](DefaultApi.md#DeleteTeamLevelDepotSpec) | **Post** /DeleteTeamLevelDepotSpec | 仓库设置-删除团队级别的分支规范
[**DeleteTeamMember**](DefaultApi.md#DeleteTeamMember) | **Post** /?action&#x3D;DeleteTeamMember | 团队成员删除
[**DeleteTestCase**](DefaultApi.md#DeleteTestCase) | **Post** /?action&#x3D;DeleteTestCase | 测试用例删除
[**DeleteTestCaseSection**](DefaultApi.md#DeleteTestCaseSection) | **Post** /?action&#x3D;DeleteTestCaseSection | 测试用例分组删除
[**DeleteTestRun**](DefaultApi.md#DeleteTestRun) | **Post** /?action&#x3D;DeleteTestRun | 测试计划删除
[**DeleteUserGroupByIds**](DefaultApi.md#DeleteUserGroupByIds) | **Post** /?action&#x3D;DeleteUserGroupByIds | 用户组删除
[**DeleteUserGroupUsers**](DefaultApi.md#DeleteUserGroupUsers) | **Post** /?action&#x3D;DeleteUserGroupUsers | 用户组删除用户
[**DescribeAgentSecret**](DefaultApi.md#DescribeAgentSecret) | **Post** /?action&#x3D;DescribeAgentSecret | 堡垒机安装 Secret
[**DescribeAllMergeRequestNotes**](DefaultApi.md#DescribeAllMergeRequestNotes) | **Post** /DescribeAllMergeRequestNotes | 合并请求-获取所有合并请求行评论和改动文件diff行评论
[**DescribeAllProjectLabels**](DefaultApi.md#DescribeAllProjectLabels) | **Post** /?action&#x3D;DescribeAllProjectLabels | 项目标签查询
[**DescribeAllProjectsIssueWorkLogList**](DefaultApi.md#DescribeAllProjectsIssueWorkLogList) | **Post** /?action&#x3D;DescribeAllProjectsIssueWorkLogList | 工时日志列表查询
[**DescribeArtifactChecksums**](DefaultApi.md#DescribeArtifactChecksums) | **Post** /?action&#x3D;DescribeArtifactChecksums | 制品Checksum列表查询
[**DescribeArtifactCredit**](DefaultApi.md#DescribeArtifactCredit) | **Post** /?action&#x3D;DescribeArtifactCredit | 查询制品授信清单详情
[**DescribeArtifactCreditList**](DefaultApi.md#DescribeArtifactCreditList) | **Post** /?action&#x3D;DescribeArtifactCreditList | 制品授信清单列表查询
[**DescribeArtifactFileDownloadUrl**](DefaultApi.md#DescribeArtifactFileDownloadUrl) | **Post** /?action&#x3D;DescribeArtifactFileDownloadUrl | 制品文件临时下载链接获取
[**DescribeArtifactPackageList**](DefaultApi.md#DescribeArtifactPackageList) | **Post** /?action&#x3D;DescribeArtifactPackageList | 制品包（镜像）列表查询
[**DescribeArtifactProperties**](DefaultApi.md#DescribeArtifactProperties) | **Post** /?action&#x3D;DescribeArtifactProperties | 制品属性列表查询
[**DescribeArtifactRepositoryFileList**](DefaultApi.md#DescribeArtifactRepositoryFileList) | **Post** /?action&#x3D;DescribeArtifactRepositoryFileList | 制品仓库下可下载的文件列表获取
[**DescribeArtifactRepositoryList**](DefaultApi.md#DescribeArtifactRepositoryList) | **Post** /?action&#x3D;DescribeArtifactRepositoryList | 制品仓库列表查询
[**DescribeArtifactVersionFileList**](DefaultApi.md#DescribeArtifactVersionFileList) | **Post** /?action&#x3D;DescribeArtifactVersionFileList | 制品版本可下载的文件列表获取
[**DescribeArtifactVersionList**](DefaultApi.md#DescribeArtifactVersionList) | **Post** /?action&#x3D;DescribeArtifactVersionList | 制品版本列表查询
[**DescribeAvailablePoliciesOnResource**](DefaultApi.md#DescribeAvailablePoliciesOnResource) | **Post** /?action&#x3D;DescribeAvailablePoliciesOnResource | 权限组列表查询（指定资源）
[**DescribeBlockIssueList**](DefaultApi.md#DescribeBlockIssueList) | **Post** /?action&#x3D;DescribeBlockIssueList | 后置事项查询
[**DescribeBlockedByIssueList**](DefaultApi.md#DescribeBlockedByIssueList) | **Post** /?action&#x3D;DescribeBlockedByIssueList | 前置事项查询
[**DescribeBranchProtection**](DefaultApi.md#DescribeBranchProtection) | **Post** /DescribeBranchProtection | 仓库设置-查询单个保护分支规则
[**DescribeBranchProtectionMembers**](DefaultApi.md#DescribeBranchProtectionMembers) | **Post** /DescribeBranchProtectionMembers | 仓库设置-查询保护分支规则下所有管理员信息
[**DescribeBranchProtections**](DefaultApi.md#DescribeBranchProtections) | **Post** /DescribeBranchProtections | 仓库设置-查询仓库保护分支规则集合
[**DescribeCanMerge**](DefaultApi.md#DescribeCanMerge) | **Post** /?action&#x3D;DescribeCanMerge | 合并请求-查看两个分支是否可以合并
[**DescribeCdAgentMachines**](DefaultApi.md#DescribeCdAgentMachines) | **Post** /?action&#x3D;DescribeCdAgentMachines | CD 堡垒机列表获取
[**DescribeCdApplication**](DefaultApi.md#DescribeCdApplication) | **Post** /?action&#x3D;DescribeCdApplication | CD 应用详情获取
[**DescribeCdApplications**](DefaultApi.md#DescribeCdApplications) | **Post** /?action&#x3D;DescribeCdApplications | CD 应用列表获取
[**DescribeCdApplicationsByProject**](DefaultApi.md#DescribeCdApplicationsByProject) | **Post** /?action&#x3D;DescribeCdApplicationsByProject | 关联应用列表获取（指定项目名）
[**DescribeCdCloudAccounts**](DefaultApi.md#DescribeCdCloudAccounts) | **Post** /?action&#x3D;DescribeCdCloudAccounts | CD 云账号列表获取
[**DescribeCdDeployCountByApplications**](DefaultApi.md#DescribeCdDeployCountByApplications) | **Post** /?action&#x3D;DescribeCdDeployCountByApplications | 发布次数-根据应用名列表获取
[**DescribeCdDeployCountByProject**](DefaultApi.md#DescribeCdDeployCountByProject) | **Post** /?action&#x3D;DescribeCdDeployCountByProject | 关联应用的发布次数获取（指定项目名）
[**DescribeCdDeployTimeByApplications**](DefaultApi.md#DescribeCdDeployTimeByApplications) | **Post** /?action&#x3D;DescribeCdDeployTimeByApplications | 发布时长-根据应用名列表获取
[**DescribeCdDeployTimeByProject**](DefaultApi.md#DescribeCdDeployTimeByProject) | **Post** /?action&#x3D;DescribeCdDeployTimeByProject | 关联应用的发布时长-根据项目名获取
[**DescribeCdDeployTrendByApplications**](DefaultApi.md#DescribeCdDeployTrendByApplications) | **Post** /?action&#x3D;DescribeCdDeployTrendByApplications | 发布趋势-根据应用名列表获取
[**DescribeCdDeployTrendByProject**](DefaultApi.md#DescribeCdDeployTrendByProject) | **Post** /?action&#x3D;DescribeCdDeployTrendByProject | 关联应用的发布趋势-根据项目名获取
[**DescribeCdHostServerGroup**](DefaultApi.md#DescribeCdHostServerGroup) | **Post** /?action&#x3D;DescribeCdHostServerGroup | CD 主机组获取
[**DescribeCdHostServerGroups**](DefaultApi.md#DescribeCdHostServerGroups) | **Post** /?action&#x3D;DescribeCdHostServerGroups | CD 主机组列表获取
[**DescribeCdPipeline**](DefaultApi.md#DescribeCdPipeline) | **Post** /?action&#x3D;DescribeCdPipeline | CD 部署流程执行记录获取
[**DescribeCdPipelineConfig**](DefaultApi.md#DescribeCdPipelineConfig) | **Post** /?action&#x3D;DescribeCdPipelineConfig | CD 部署流程配置-根据名称获取
[**DescribeCdPipelineConfigs**](DefaultApi.md#DescribeCdPipelineConfigs) | **Post** /?action&#x3D;DescribeCdPipelineConfigs | CD 应用下的所有部署流程配置获取
[**DescribeCdTask**](DefaultApi.md#DescribeCdTask) | **Post** /?action&#x3D;DescribeCdTask | CD 任务执行记录获取
[**DescribeCodeSearch**](DefaultApi.md#DescribeCodeSearch) | **Post** /DescribeCodeSearch | 仓库信息-查询代码片段详细列表
[**DescribeCodingCIBuild**](DefaultApi.md#DescribeCodingCIBuild) | **Post** /?action&#x3D;DescribeCodingCIBuild | 构建记录详情查询
[**DescribeCodingCIBuildArtifacts**](DefaultApi.md#DescribeCodingCIBuildArtifacts) | **Post** /?action&#x3D;DescribeCodingCIBuildArtifacts | 构建任务制品查询
[**DescribeCodingCIBuildEnvs**](DefaultApi.md#DescribeCodingCIBuildEnvs) | **Post** /?action&#x3D;DescribeCodingCIBuildEnvs | 构建计划环境变量获取
[**DescribeCodingCIBuildHtmlReports**](DefaultApi.md#DescribeCodingCIBuildHtmlReports) | **Post** /?action&#x3D;DescribeCodingCIBuildHtmlReports | 构建任务网页报告查询
[**DescribeCodingCIBuildLog**](DefaultApi.md#DescribeCodingCIBuildLog) | **Post** /?action&#x3D;DescribeCodingCIBuildLog | 构建日志获取
[**DescribeCodingCIBuildLogRaw**](DefaultApi.md#DescribeCodingCIBuildLogRaw) | **Post** /?action&#x3D;DescribeCodingCIBuildLogRaw | 构建完整日志查询（原始日志 Raw）
[**DescribeCodingCIBuildMetrics**](DefaultApi.md#DescribeCodingCIBuildMetrics) | **Post** /?action&#x3D;DescribeCodingCIBuildMetrics | 构建计划度量查询
[**DescribeCodingCIBuildStage**](DefaultApi.md#DescribeCodingCIBuildStage) | **Post** /?action&#x3D;DescribeCodingCIBuildStage | 构建任务阶段获取
[**DescribeCodingCIBuildStatistics**](DefaultApi.md#DescribeCodingCIBuildStatistics) | **Post** /?action&#x3D;DescribeCodingCIBuildStatistics | 构建任务统计
[**DescribeCodingCIBuildStep**](DefaultApi.md#DescribeCodingCIBuildStep) | **Post** /?action&#x3D;DescribeCodingCIBuildStep | 构建任务指定阶段的步骤获取
[**DescribeCodingCIBuildStepLog**](DefaultApi.md#DescribeCodingCIBuildStepLog) | **Post** /?action&#x3D;DescribeCodingCIBuildStepLog | 构建步骤日志获取
[**DescribeCodingCIBuilds**](DefaultApi.md#DescribeCodingCIBuilds) | **Post** /?action&#x3D;DescribeCodingCIBuilds | 构建计划的构建列表获取
[**DescribeCodingCIJob**](DefaultApi.md#DescribeCodingCIJob) | **Post** /?action&#x3D;DescribeCodingCIJob | 构建计划详情获取
[**DescribeCodingCIJobs**](DefaultApi.md#DescribeCodingCIJobs) | **Post** /?action&#x3D;DescribeCodingCIJobs | 构建计划查询（通过项目ID）
[**DescribeCodingCurrentUser**](DefaultApi.md#DescribeCodingCurrentUser) | **Post** /?action&#x3D;DescribeCodingCurrentUser | 当前用户信息查询
[**DescribeCodingProjects**](DefaultApi.md#DescribeCodingProjects) | **Post** /?action&#x3D;DescribeCodingProjects | 项目列表查询
[**DescribeCommitRefs**](DefaultApi.md#DescribeCommitRefs) | **Post** /DescribeCommitRefs | Git提交-查询commit的ref信息
[**DescribeCommitsBetweenCommitAndCommit**](DefaultApi.md#DescribeCommitsBetweenCommitAndCommit) | **Post** /DescribeCommitsBetweenCommitAndCommit | Git提交-查询两个请求之间的请求列表（source target顺序正常）
[**DescribeConfigTemplateList**](DefaultApi.md#DescribeConfigTemplateList) | **Post** /DescribeConfigTemplateList | 配置方案获取
[**DescribeDepartment**](DefaultApi.md#DescribeDepartment) | **Post** /?action&#x3D;DescribeDepartment | 部门详情查询
[**DescribeDepartmentMembers**](DefaultApi.md#DescribeDepartmentMembers) | **Post** /?action&#x3D;DescribeDepartmentMembers | 部门成员列表查询
[**DescribeDepotByNameInfo**](DefaultApi.md#DescribeDepotByNameInfo) | **Post** /DescribeDepotByNameInfo | 仓库信息-查询项目下所有的仓库信息列表
[**DescribeDepotDefaultBranch**](DefaultApi.md#DescribeDepotDefaultBranch) | **Post** /DescribeDepotDefaultBranch | 仓库分支-查询仓库的默认分支
[**DescribeDepotFilePushRules**](DefaultApi.md#DescribeDepotFilePushRules) | **Post** /DescribeDepotFilePushRules | 仓库设置-查询git仓库文件推送规则
[**DescribeDepotMergeRequests**](DefaultApi.md#DescribeDepotMergeRequests) | **Post** /DescribeDepotMergeRequests | 合并请求-查询仓库合并请求列表
[**DescribeDepotPushSetting**](DefaultApi.md#DescribeDepotPushSetting) | **Post** /DescribeDepotPushSetting | 仓库设置-查询仓库推送设置
[**DescribeDepotSpecDetail**](DefaultApi.md#DescribeDepotSpecDetail) | **Post** /DescribeDepotSpecDetail | 仓库设置-查询仓库规范详情
[**DescribeDepotSpecs**](DefaultApi.md#DescribeDepotSpecs) | **Post** /DescribeDepotSpecs | 仓库设置-查询仓库规范列表
[**DescribeDifferentBetween2Commits**](DefaultApi.md#DescribeDifferentBetween2Commits) | **Post** /?action&#x3D;DescribeDifferentBetween2Commits | Git提交-两次提交之间的文件差异（source target顺序正常）
[**DescribeDifferentBetweenTwoCommits**](DefaultApi.md#DescribeDifferentBetweenTwoCommits) | **Post** /DescribeDifferentBetweenTwoCommits | Git提交-获取两次commit之间的文件差异详情(废弃，source target顺序不一致)
[**DescribeGitBlameInfo**](DefaultApi.md#DescribeGitBlameInfo) | **Post** /DescribeGitBlameInfo | Git提交-获取指定commit下某文件指定代码行的最后一次提交
[**DescribeGitBlob**](DefaultApi.md#DescribeGitBlob) | **Post** /DescribeGitBlob | Git文件-查询GitBlob
[**DescribeGitBlobRaw**](DefaultApi.md#DescribeGitBlobRaw) | **Post** /DescribeGitBlobRaw | Git文件-查询Git Blob raw信息
[**DescribeGitBranch**](DefaultApi.md#DescribeGitBranch) | **Post** /DescribeGitBranch | 仓库分支-查询代码仓库单个分支
[**DescribeGitBranchList**](DefaultApi.md#DescribeGitBranchList) | **Post** /DescribeGitBranchList | 仓库分支-查询仓库分支列表
[**DescribeGitBranches**](DefaultApi.md#DescribeGitBranches) | **Post** /DescribeGitBranches | 仓库分支-查询仓库下所有分支列表
[**DescribeGitBranchesBySha**](DefaultApi.md#DescribeGitBranchesBySha) | **Post** /DescribeGitBranchesBySha | 仓库分支-根据sha值查询所在分支
[**DescribeGitCommitComments**](DefaultApi.md#DescribeGitCommitComments) | **Post** /DescribeGitCommitComments | Git提交-获取commit评论
[**DescribeGitCommitDiff**](DefaultApi.md#DescribeGitCommitDiff) | **Post** /DescribeGitCommitDiff | Git提交-查询某次提交的diff信息
[**DescribeGitCommitFilePathList**](DefaultApi.md#DescribeGitCommitFilePathList) | **Post** /DescribeGitCommitFilePathList | Git提交-查询仓库某次提交改动的文件路径列表
[**DescribeGitCommitInfo**](DefaultApi.md#DescribeGitCommitInfo) | **Post** /DescribeGitCommitInfo | Git提交-查询单个请求详情信息
[**DescribeGitCommitInfos**](DefaultApi.md#DescribeGitCommitInfos) | **Post** /DescribeGitCommitInfos | Git提交-查询仓库分支下提交列表
[**DescribeGitCommitNote**](DefaultApi.md#DescribeGitCommitNote) | **Post** /DescribeGitCommitNote | Git提交-获取提交注释
[**DescribeGitCommitStatus**](DefaultApi.md#DescribeGitCommitStatus) | **Post** /DescribeGitCommitStatus | Git提交-查询提交对应的流水线状态
[**DescribeGitCommitsInPage**](DefaultApi.md#DescribeGitCommitsInPage) | **Post** /DescribeGitCommitsInPage | Git提交-查询仓库分支下提交列表
[**DescribeGitContributors**](DefaultApi.md#DescribeGitContributors) | **Post** /DescribeGitContributors | 仓库信息-查询git仓库的贡献者
[**DescribeGitDepot**](DefaultApi.md#DescribeGitDepot) | **Post** /DescribeGitDepot | 仓库信息-根据代码仓库id获取代码仓库信息
[**DescribeGitDepotDeployKeys**](DefaultApi.md#DescribeGitDepotDeployKeys) | **Post** /DescribeGitDepotDeployKeys | 仓库设置-查询某仓库下的部署公钥列表
[**DescribeGitFile**](DefaultApi.md#DescribeGitFile) | **Post** /DescribeGitFile | Git文件-获取文件详情
[**DescribeGitFileContent**](DefaultApi.md#DescribeGitFileContent) | **Post** /DescribeGitFileContent | Git提交-查询某次提交某文件的内容
[**DescribeGitFileStat**](DefaultApi.md#DescribeGitFileStat) | **Post** /DescribeGitFileStat | Git文件-检查仓库文件是否存在
[**DescribeGitFiles**](DefaultApi.md#DescribeGitFiles) | **Post** /DescribeGitFiles | Git文件-查询仓库目录下文件和文件夹名字
[**DescribeGitMergeBase**](DefaultApi.md#DescribeGitMergeBase) | **Post** /DescribeGitMergeBase | 仓库分支-查询两个分支的公共祖先
[**DescribeGitMergeRequestDiffDetail**](DefaultApi.md#DescribeGitMergeRequestDiffDetail) | **Post** /DescribeGitMergeRequestDiffDetail | 合并请求-查询合并请求文件的 diff 详情
[**DescribeGitMergeRequestDiffs**](DefaultApi.md#DescribeGitMergeRequestDiffs) | **Post** /DescribeGitMergeRequestDiffs | 合并请求-查询合并请求diff信息的文件列表
[**DescribeGitMergeRequestParticipants**](DefaultApi.md#DescribeGitMergeRequestParticipants) | **Post** /DescribeGitMergeRequestParticipants | 合并请求-获取合并请求的参与者
[**DescribeGitMergeRequestsBySha**](DefaultApi.md#DescribeGitMergeRequestsBySha) | **Post** /DescribeGitMergeRequestsBySha | 合并请求-查询含有某次提交的合并请求
[**DescribeGitProjectDeployKeys**](DefaultApi.md#DescribeGitProjectDeployKeys) | **Post** /DescribeGitProjectDeployKeys | 仓库设置-查询某项目下的部署公钥列表
[**DescribeGitProtectedTags**](DefaultApi.md#DescribeGitProtectedTags) | **Post** /DescribeGitProtectedTags | 标签信息-查询受保护的标签列表
[**DescribeGitProtectedTagsByRule**](DefaultApi.md#DescribeGitProtectedTagsByRule) | **Post** /DescribeGitProtectedTagsByRule | 标签信息-根据标签保护规则查询受保护的标签列表
[**DescribeGitRef**](DefaultApi.md#DescribeGitRef) | **Post** /DescribeGitRef | 仓库分支-获取分支或标签信息
[**DescribeGitRefsBySha**](DefaultApi.md#DescribeGitRefsBySha) | **Post** /DescribeGitRefsBySha | Git提交-查询含有某次提交的标签或分支列表
[**DescribeGitReleaseDetail**](DefaultApi.md#DescribeGitReleaseDetail) | **Post** /DescribeGitReleaseDetail | 版本信息-查询仓库的版本详情
[**DescribeGitReleases**](DefaultApi.md#DescribeGitReleases) | **Post** /DescribeGitReleases | 版本信息-查询仓库的版本列表
[**DescribeGitTag**](DefaultApi.md#DescribeGitTag) | **Post** /DescribeGitTag | 标签信息-查询单个tag
[**DescribeGitTags**](DefaultApi.md#DescribeGitTags) | **Post** /DescribeGitTags | 标签信息-查询当前仓库下所有tags
[**DescribeGitTagsByBranch**](DefaultApi.md#DescribeGitTagsByBranch) | **Post** /DescribeGitTagsByBranch | 标签信息-根据分支获取标签列表
[**DescribeGitTagsBySha**](DefaultApi.md#DescribeGitTagsBySha) | **Post** /DescribeGitTagsBySha | 标签信息-查询含有某次提交的标签列表
[**DescribeGitTree**](DefaultApi.md#DescribeGitTree) | **Post** /DescribeGitTree | 仓库信息-查询git仓库的树
[**DescribeGrantObjectsOnResource**](DefaultApi.md#DescribeGrantObjectsOnResource) | **Post** /?action&#x3D;DescribeGrantObjectsOnResource | 授权对象列表分页查询
[**DescribeGrantUsersOnResource**](DefaultApi.md#DescribeGrantUsersOnResource) | **Post** /?action&#x3D;DescribeGrantUsersOnResource | 授权用户列表分页查询
[**DescribeHostServerInstance**](DefaultApi.md#DescribeHostServerInstance) | **Post** /?action&#x3D;DescribeHostServerInstance | CD 主机组部署详情获取
[**DescribeIssue**](DefaultApi.md#DescribeIssue) | **Post** /DescribeIssue | 事项详情查询
[**DescribeIssueAttachmentPreSignedUrl**](DefaultApi.md#DescribeIssueAttachmentPreSignedUrl) | **Post** /?action&#x3D;DescribeIssueAttachmentPreSignedUrl | 预签名信息获取
[**DescribeIssueCommentList**](DefaultApi.md#DescribeIssueCommentList) | **Post** /?action&#x3D;DescribeIssueCommentList | 事项评论列表查询
[**DescribeIssueCustomFieldLogList**](DefaultApi.md#DescribeIssueCustomFieldLogList) | **Post** /DescribeIssueCustomFieldLogList | 事项的自定义属性变更日志查询
[**DescribeIssueFileUrl**](DefaultApi.md#DescribeIssueFileUrl) | **Post** /?action&#x3D;DescribeIssueFileUrl | 事项附件的下载地址查询
[**DescribeIssueFilterList**](DefaultApi.md#DescribeIssueFilterList) | **Post** /?action&#x3D;DescribeIssueFilterList | 事项筛选器列表查询
[**DescribeIssueList**](DefaultApi.md#DescribeIssueList) | **Post** /?action&#x3D;DescribeIssueList | 事项列表查询
[**DescribeIssueListWithPage**](DefaultApi.md#DescribeIssueListWithPage) | **Post** /DescribeIssueListWithPage | 事项查询,返回分页信息
[**DescribeIssueModuleList**](DefaultApi.md#DescribeIssueModuleList) | **Post** /DescribeIssueModuleList | 事项模块列表查询
[**DescribeIssueReferenceResources**](DefaultApi.md#DescribeIssueReferenceResources) | **Post** /DescribeIssueReferenceResources | 事项的引用资源列表查询
[**DescribeIssueRelatedWorkItemList**](DefaultApi.md#DescribeIssueRelatedWorkItemList) | **Post** /DescribeIssueRelatedWorkItemList | 事项关联的项目集中的工作项查询
[**DescribeIssueReleaseList**](DefaultApi.md#DescribeIssueReleaseList) | **Post** /DescribeIssueReleaseList | 事项加入的版本查询
[**DescribeIssueStatusChangeLogList**](DefaultApi.md#DescribeIssueStatusChangeLogList) | **Post** /DescribeIssueStatusChangeLogList | 事项的状态变更记录查询
[**DescribeIssueWorkLogList**](DefaultApi.md#DescribeIssueWorkLogList) | **Post** /DescribeIssueWorkLogList | 事项的工时日志查询
[**DescribeIteration**](DefaultApi.md#DescribeIteration) | **Post** /?action&#x3D;DescribeIteration | 迭代详情查询
[**DescribeIterationList**](DefaultApi.md#DescribeIterationList) | **Post** /?action&#x3D;DescribeIterationList | 迭代列表获取
[**DescribeMemberSshKey**](DefaultApi.md#DescribeMemberSshKey) | **Post** /DescribeMemberSshKey | 仓库设置-获取团队成员的SSH公钥列表
[**DescribeMergeReqCommits**](DefaultApi.md#DescribeMergeReqCommits) | **Post** /DescribeMergeReqCommits | 合并请求-查询合并请求列表
[**DescribeMergeReqInfo**](DefaultApi.md#DescribeMergeReqInfo) | **Post** /DescribeMergeReqInfo | 合并请求-查询合并请求详情
[**DescribeMergeRequest**](DefaultApi.md#DescribeMergeRequest) | **Post** /DescribeMergeRequest | 合并请求-查询合并请求详情信息
[**DescribeMergeRequestFileDiff**](DefaultApi.md#DescribeMergeRequestFileDiff) | **Post** /DescribeMergeRequestFileDiff | 合并请求-获取合并请求文件修改记录
[**DescribeMergeRequestLog**](DefaultApi.md#DescribeMergeRequestLog) | **Post** /DescribeMergeRequestLog | 合并请求-查询合并请求操作记录
[**DescribeMergeRequestReviewers**](DefaultApi.md#DescribeMergeRequestReviewers) | **Post** /DescribeMergeRequestReviewers | 合并请求-获取合并请求的评审者
[**DescribeMyDepots**](DefaultApi.md#DescribeMyDepots) | **Post** /DescribeMyDepots | 仓库信息-获取当前用户拥有读权限的仓库列表
[**DescribeNotesByCommits**](DefaultApi.md#DescribeNotesByCommits) | **Post** /DescribeNotesByCommits | 仓库信息-获取提交的note信息
[**DescribeOneProject**](DefaultApi.md#DescribeOneProject) | **Post** /?action&#x3D;DescribeOneProject | 单个项目查询
[**DescribePersonalExternalDepots**](DefaultApi.md#DescribePersonalExternalDepots) | **Post** /?action&#x3D;DescribePersonalExternalDepots | 个人外部仓库获取
[**DescribePinyin**](DefaultApi.md#DescribePinyin) | **Post** /?action&#x3D;DescribePinyin | 汉字转拼音
[**DescribePoliciesOnResourceType**](DefaultApi.md#DescribePoliciesOnResourceType) | **Post** /?action&#x3D;DescribePoliciesOnResourceType | 权限组列表查询（指定资源类型）
[**DescribePolicy**](DefaultApi.md#DescribePolicy) | **Post** /?action&#x3D;DescribePolicy | 权限组详情获取
[**DescribePreSignUploadUrl**](DefaultApi.md#DescribePreSignUploadUrl) | **Post** /?action&#x3D;DescribePreSignUploadUrl | 预签名URL获取
[**DescribePredicatePolicyOnResource**](DefaultApi.md#DescribePredicatePolicyOnResource) | **Post** /?action&#x3D;DescribePredicatePolicyOnResource | 资源权限判定模式获取
[**DescribeProgramProjects**](DefaultApi.md#DescribeProgramProjects) | **Post** /?action&#x3D;DescribeProgramProjects | 项目集下项目列表查询
[**DescribePrograms**](DefaultApi.md#DescribePrograms) | **Post** /?action&#x3D;DescribePrograms | 项目集列表查询
[**DescribeProjectAnnouncement**](DefaultApi.md#DescribeProjectAnnouncement) | **Post** /?action&#x3D;DescribeProjectAnnouncement | 项目公告查询
[**DescribeProjectAnnouncements**](DefaultApi.md#DescribeProjectAnnouncements) | **Post** /?action&#x3D;DescribeProjectAnnouncements | 项目公告列表查询
[**DescribeProjectByName**](DefaultApi.md#DescribeProjectByName) | **Post** /?action&#x3D;DescribeProjectByName | 项目查询(通过项目名称)
[**DescribeProjectCredentials**](DefaultApi.md#DescribeProjectCredentials) | **Post** /?action&#x3D;DescribeProjectCredentials | 项目凭据列表查询
[**DescribeProjectDepotBranches**](DefaultApi.md#DescribeProjectDepotBranches) | **Post** /?action&#x3D;DescribeProjectDepotBranches | 仓库分支列表获取
[**DescribeProjectDepotCommits**](DefaultApi.md#DescribeProjectDepotCommits) | **Post** /?action&#x3D;DescribeProjectDepotCommits | 分支下的提交列表获取
[**DescribeProjectDepotInfoList**](DefaultApi.md#DescribeProjectDepotInfoList) | **Post** /DescribeProjectDepotInfoList | 仓库信息-查询项目下所有的仓库信息列表
[**DescribeProjectDepotTags**](DefaultApi.md#DescribeProjectDepotTags) | **Post** /?action&#x3D;DescribeProjectDepotTags | 仓库的标签列表获取
[**DescribeProjectDepots**](DefaultApi.md#DescribeProjectDepots) | **Post** /?action&#x3D;DescribeProjectDepots | 项目仓库列表获取
[**DescribeProjectIssueFieldList**](DefaultApi.md#DescribeProjectIssueFieldList) | **Post** /?action&#x3D;DescribeProjectIssueFieldList | 具体事项类型的属性列表查询
[**DescribeProjectIssueStatusList**](DefaultApi.md#DescribeProjectIssueStatusList) | **Post** /?action&#x3D;DescribeProjectIssueStatusList | 具体事项类型的状态列表查询
[**DescribeProjectIssueTypeList**](DefaultApi.md#DescribeProjectIssueTypeList) | **Post** /?action&#x3D;DescribeProjectIssueTypeList | 项目事项类型列表查询
[**DescribeProjectLabels**](DefaultApi.md#DescribeProjectLabels) | **Post** /?action&#x3D;DescribeProjectLabels | 项目列表查询-指定项目标签
[**DescribeProjectMemberPrincipals**](DefaultApi.md#DescribeProjectMemberPrincipals) | **Post** /?action&#x3D;DescribeProjectMemberPrincipals | 项目成员主体查询(包含用户组、部门、成员)
[**DescribeProjectMembers**](DefaultApi.md#DescribeProjectMembers) | **Post** /?action&#x3D;DescribeProjectMembers | 项目成员列表查询
[**DescribeProjectMergeRequests**](DefaultApi.md#DescribeProjectMergeRequests) | **Post** /DescribeProjectMergeRequests | 合并请求-获取项目下的合并请求列表
[**DescribeProjectRoles**](DefaultApi.md#DescribeProjectRoles) | **Post** /?action&#x3D;DescribeProjectRoles | 项目用户组查询
[**DescribeProjectsByFeature**](DefaultApi.md#DescribeProjectsByFeature) | **Post** /?action&#x3D;DescribeProjectsByFeature | 项目查询（通过一级菜单名）
[**DescribeProtectedBranch**](DefaultApi.md#DescribeProtectedBranch) | **Post** /DescribeProtectedBranch | 仓库设置-查询保护分支详情
[**DescribeProtectedBranchMembers**](DefaultApi.md#DescribeProtectedBranchMembers) | **Post** /DescribeProtectedBranchMembers | 仓库设置-查询保护分支成员
[**DescribeProtectedBranches**](DefaultApi.md#DescribeProtectedBranches) | **Post** /DescribeProtectedBranches | 仓库设置-查询保护分支列表
[**DescribeRelatedCaseList**](DefaultApi.md#DescribeRelatedCaseList) | **Post** /DescribeRelatedCaseList | 事项关联的测试用例查询
[**DescribeRelease**](DefaultApi.md#DescribeRelease) | **Post** /DescribeRelease | 版本详情查询
[**DescribeReleaseIssueList**](DefaultApi.md#DescribeReleaseIssueList) | **Post** /DescribeReleaseIssueList | 版本发布范围查询
[**DescribeReleaseList**](DefaultApi.md#DescribeReleaseList) | **Post** /DescribeReleaseList | 版本列表查询
[**DescribeReport**](DefaultApi.md#DescribeReport) | **Post** /?action&#x3D;DescribeReport | 测试报告详情
[**DescribeReportList**](DefaultApi.md#DescribeReportList) | **Post** /?action&#x3D;DescribeReportList | 测试报告列表
[**DescribeRequirementDefectRelation**](DefaultApi.md#DescribeRequirementDefectRelation) | **Post** /?action&#x3D;DescribeRequirementDefectRelation | 需求关联缺陷列表查询
[**DescribeRequirementTestCaseList**](DefaultApi.md#DescribeRequirementTestCaseList) | **Post** /?action&#x3D;DescribeRequirementTestCaseList | 需求关联的测试用例列表
[**DescribeResourceReferences**](DefaultApi.md#DescribeResourceReferences) | **Post** /?action&#x3D;DescribeResourceReferences | 资源引用的资源列表，如 开发任务中引用了多个需求，获取任务引用的需求列表
[**DescribeResourceReferencesCited**](DefaultApi.md#DescribeResourceReferencesCited) | **Post** /?action&#x3D;DescribeResourceReferencesCited | 被引用资源列表查询
[**DescribeResourceReferencesCiting**](DefaultApi.md#DescribeResourceReferencesCiting) | **Post** /?action&#x3D;DescribeResourceReferencesCiting | 引用资源列表查询
[**DescribeResourceScopeListOnPolicy**](DefaultApi.md#DescribeResourceScopeListOnPolicy) | **Post** /?action&#x3D;DescribeResourceScopeListOnPolicy | 权限组可用资源范围分页查询
[**DescribeSelfMergeRequests**](DefaultApi.md#DescribeSelfMergeRequests) | **Post** /DescribeSelfMergeRequests | 合并请求-获取自己的合并请求列表
[**DescribeSingeMergeRequestNotes**](DefaultApi.md#DescribeSingeMergeRequestNotes) | **Post** /DescribeSingeMergeRequestNotes | 合并请求-获取单个合并请求行评论和改动文件diff行评论
[**DescribeSshKey**](DefaultApi.md#DescribeSshKey) | **Post** /?action&#x3D;DescribeSshKey | 仓库设置-获取当前用户所有SSH公钥
[**DescribeSubIssueList**](DefaultApi.md#DescribeSubIssueList) | **Post** /?action&#x3D;DescribeSubIssueList | 子事项列表查询
[**DescribeTeam**](DefaultApi.md#DescribeTeam) | **Post** /?action&#x3D;DescribeTeam | 团队信息查询
[**DescribeTeamAdminMembers**](DefaultApi.md#DescribeTeamAdminMembers) | **Post** /?action&#x3D;DescribeTeamAdminMembers | 团队管理员查询
[**DescribeTeamArtifacts**](DefaultApi.md#DescribeTeamArtifacts) | **Post** /?action&#x3D;DescribeTeamArtifacts | 制品列表查询
[**DescribeTeamDepotInfoList**](DefaultApi.md#DescribeTeamDepotInfoList) | **Post** /DescribeTeamDepotInfoList | 仓库信息-获取团队下仓库列表，仅团队所有者或团队管理员可以调用。
[**DescribeTeamDepots**](DefaultApi.md#DescribeTeamDepots) | **Post** /?action&#x3D;DescribeTeamDepots | 团队下可访问的所有仓库列表获取
[**DescribeTeamIssueTypeList**](DefaultApi.md#DescribeTeamIssueTypeList) | **Post** /?action&#x3D;DescribeTeamIssueTypeList | 企业事项类型列表查询
[**DescribeTeamMember**](DefaultApi.md#DescribeTeamMember) | **Post** /?action&#x3D;DescribeTeamMember | 团队成员信息查询（通过用户 ID）
[**DescribeTeamMemberByEmail**](DefaultApi.md#DescribeTeamMemberByEmail) | **Post** /?action&#x3D;DescribeTeamMemberByEmail | 团队成员信息查询（通过用户 Email）
[**DescribeTeamMembers**](DefaultApi.md#DescribeTeamMembers) | **Post** /?action&#x3D;DescribeTeamMembers | 团队成员列表查询
[**DescribeTest**](DefaultApi.md#DescribeTest) | **Post** /?action&#x3D;DescribeTest | 测试任务详情
[**DescribeTestCase**](DefaultApi.md#DescribeTestCase) | **Post** /?action&#x3D;DescribeTestCase | 测试用例详情
[**DescribeTestCaseList**](DefaultApi.md#DescribeTestCaseList) | **Post** /?action&#x3D;DescribeTestCaseList | 测试用例列表
[**DescribeTestCaseSectionList**](DefaultApi.md#DescribeTestCaseSectionList) | **Post** /?action&#x3D;DescribeTestCaseSectionList | 测试用例分组列表
[**DescribeTestDefectList**](DefaultApi.md#DescribeTestDefectList) | **Post** /?action&#x3D;DescribeTestDefectList | 测试任务关联的缺陷列表
[**DescribeTestList**](DefaultApi.md#DescribeTestList) | **Post** /?action&#x3D;DescribeTestList | 测试任务列表
[**DescribeTestRun**](DefaultApi.md#DescribeTestRun) | **Post** /?action&#x3D;DescribeTestRun | 测试计划详情
[**DescribeTestRunList**](DefaultApi.md#DescribeTestRunList) | **Post** /?action&#x3D;DescribeTestRunList | 测试计划列表
[**DescribeUserGroups**](DefaultApi.md#DescribeUserGroups) | **Post** /?action&#x3D;DescribeUserGroups | 用户组列表分页查询
[**DescribeUserProjects**](DefaultApi.md#DescribeUserProjects) | **Post** /?action&#x3D;DescribeUserProjects | 项目列表查询（已加入的项目）
[**DescribeUsersByGroupId**](DefaultApi.md#DescribeUsersByGroupId) | **Post** /?action&#x3D;DescribeUsersByGroupId | 用户列表查询（根据用户组id分页查询）
[**DescribeUsersOnResourceAndGrantObject**](DefaultApi.md#DescribeUsersOnResourceAndGrantObject) | **Post** /?action&#x3D;DescribeUsersOnResourceAndGrantObject | 授权用户列表分页查询（指定资源）
[**DescribeWorkItemSalvage**](DefaultApi.md#DescribeWorkItemSalvage) | **Post** /DescribeWorkItemSalvage | 事项分解信息查询
[**DescribeWorkbenchIssueList**](DefaultApi.md#DescribeWorkbenchIssueList) | **Post** /DescribeWorkbenchIssueList | 用户在团队内的所有代办事项查询
[**DetachFromResource**](DefaultApi.md#DetachFromResource) | **Post** /?action&#x3D;DetachFromResource | 授权收回，只收回参数指定的授权，已有其它授权不受影响
[**DetachResourceScopeOnPolicy**](DefaultApi.md#DetachResourceScopeOnPolicy) | **Post** /?action&#x3D;DetachResourceScopeOnPolicy | 权限组可用资源删除，只删除参数指定的资源，已有的其它资源不受影响
[**ForbiddenArtifactVersion**](DefaultApi.md#ForbiddenArtifactVersion) | **Post** /?action&#x3D;ForbiddenArtifactVersion | 制品版本下载 禁止、解禁
[**ModifyArtifactCredit**](DefaultApi.md#ModifyArtifactCredit) | **Post** /?action&#x3D;ModifyArtifactCredit | 制品授信清单修改
[**ModifyArtifactProperties**](DefaultApi.md#ModifyArtifactProperties) | **Post** /?action&#x3D;ModifyArtifactProperties | 制品属性修改
[**ModifyBranchProtection**](DefaultApi.md#ModifyBranchProtection) | **Post** /ModifyBranchProtection | 仓库设置-修改保护分支规则相关信息
[**ModifyBranchProtectionMemberPermission**](DefaultApi.md#ModifyBranchProtectionMemberPermission) | **Post** /ModifyBranchProtectionMemberPermission | 仓库设置-更改保护分支管理员权限
[**ModifyCdCloudAccount**](DefaultApi.md#ModifyCdCloudAccount) | **Post** /?action&#x3D;ModifyCdCloudAccount | CD 云账号更新
[**ModifyCdHostServerGroup**](DefaultApi.md#ModifyCdHostServerGroup) | **Post** /?action&#x3D;ModifyCdHostServerGroup | CD 主机组修改
[**ModifyCdPipeline**](DefaultApi.md#ModifyCdPipeline) | **Post** /?action&#x3D;ModifyCdPipeline | CD 部署流程修改
[**ModifyChooseDepotSpec**](DefaultApi.md#ModifyChooseDepotSpec) | **Post** /ModifyChooseDepotSpec | 仓库设置-使用、取消使用仓库规范
[**ModifyCloseMR**](DefaultApi.md#ModifyCloseMR) | **Post** /ModifyCloseMR | 合并请求-关闭合并请求
[**ModifyCodingCIAgentEnable**](DefaultApi.md#ModifyCodingCIAgentEnable) | **Post** /?action&#x3D;ModifyCodingCIAgentEnable | 自定义构建节点启用、禁用
[**ModifyCodingCIJob**](DefaultApi.md#ModifyCodingCIJob) | **Post** /?action&#x3D;ModifyCodingCIJob | 构建计划修改
[**ModifyDefaultBranch**](DefaultApi.md#ModifyDefaultBranch) | **Post** /ModifyDefaultBranch | 仓库设置-修改仓库默认分支
[**ModifyDefectRelatedRequirement**](DefaultApi.md#ModifyDefectRelatedRequirement) | **Post** /?action&#x3D;ModifyDefectRelatedRequirement | 缺陷所属的需求修改
[**ModifyDepartment**](DefaultApi.md#ModifyDepartment) | **Post** /?action&#x3D;ModifyDepartment | 部门信息修改
[**ModifyDepartmentAssignee**](DefaultApi.md#ModifyDepartmentAssignee) | **Post** /?action&#x3D;ModifyDepartmentAssignee | 部门负责人管理
[**ModifyDepartmentMember**](DefaultApi.md#ModifyDepartmentMember) | **Post** /?action&#x3D;ModifyDepartmentMember | 部门成员管理
[**ModifyDepotDescription**](DefaultApi.md#ModifyDepotDescription) | **Post** /ModifyDepotDescription | 仓库信息-修改仓库描述
[**ModifyDepotFilePushRule**](DefaultApi.md#ModifyDepotFilePushRule) | **Post** /ModifyDepotFilePushRule | 仓库设置-修改git仓库文件推送规则
[**ModifyDepotFilePushRuleDenyPrivilege**](DefaultApi.md#ModifyDepotFilePushRuleDenyPrivilege) | **Post** /ModifyDepotFilePushRuleDenyPrivilege | 仓库设置-修改 git 仓库特权者文件推送权限
[**ModifyDepotLevelDepotSpec**](DefaultApi.md#ModifyDepotLevelDepotSpec) | **Post** /ModifyDepotLevelDepotSpec | 仓库设置-修改、新增仓库级别的仓库规范
[**ModifyDepotName**](DefaultApi.md#ModifyDepotName) | **Post** /ModifyDepotName | 仓库信息-修改仓库名称
[**ModifyDepotPushSetting**](DefaultApi.md#ModifyDepotPushSetting) | **Post** /ModifyDepotPushSetting | 仓库设置-修改仓库推送设置
[**ModifyDepotQuota**](DefaultApi.md#ModifyDepotQuota) | **Post** /ModifyDepotQuota | 仓库信息-修改仓库容量
[**ModifyDepotSettings**](DefaultApi.md#ModifyDepotSettings) | **Post** /ModifyDepotSettings | 仓库设置-修改仓库设置
[**ModifyDepotSharedSetting**](DefaultApi.md#ModifyDepotSharedSetting) | **Post** /ModifyDepotSharedSetting | 仓库设置-修改仓库是否开源设置
[**ModifyGitCherryPick**](DefaultApi.md#ModifyGitCherryPick) | **Post** /ModifyGitCherryPick | Git提交-将某次提交cherry-pick到指定分支
[**ModifyGitCommitRevert**](DefaultApi.md#ModifyGitCommitRevert) | **Post** /ModifyGitCommitRevert | Git提交-还原某次提交
[**ModifyGitCommitStatus**](DefaultApi.md#ModifyGitCommitStatus) | **Post** /ModifyGitCommitStatus | Git提交-修改提交对应的流水线状态
[**ModifyGitDepotArchive**](DefaultApi.md#ModifyGitDepotArchive) | **Post** /ModifyGitDepotArchive | 仓库设置-仓库归档
[**ModifyGitDepotUnarchive**](DefaultApi.md#ModifyGitDepotUnarchive) | **Post** /ModifyGitDepotUnarchive | 仓库设置-仓库解除归档
[**ModifyGitFiles**](DefaultApi.md#ModifyGitFiles) | **Post** /ModifyGitFiles | Git提交-修改仓库某文件
[**ModifyGitMergeBranch**](DefaultApi.md#ModifyGitMergeBranch) | **Post** /ModifyGitMergeBranch | 合并请求-将源分支的改动合并到目标分支
[**ModifyGitMergeRequest**](DefaultApi.md#ModifyGitMergeRequest) | **Post** /ModifyGitMergeRequest | 合并请求-修改合并请求的标题和描述信息
[**ModifyGitMergeRequestRebase**](DefaultApi.md#ModifyGitMergeRequestRebase) | **Post** /ModifyGitMergeRequestRebase | 合并请求-合并请求中的源分支进行rebase操作
[**ModifyGitRebase**](DefaultApi.md#ModifyGitRebase) | **Post** /ModifyGitRebase | 仓库信息-git变基操作
[**ModifyGitRelease**](DefaultApi.md#ModifyGitRelease) | **Post** /ModifyGitRelease | 版本信息-修改仓库版本信息
[**ModifyGitTransfer**](DefaultApi.md#ModifyGitTransfer) | **Post** /ModifyGitTransfer | 仓库信息-仓库转移至同团队下的其他项目中
[**ModifyIssue**](DefaultApi.md#ModifyIssue) | **Post** /ModifyIssue | 事项修改
[**ModifyIssueComment**](DefaultApi.md#ModifyIssueComment) | **Post** /?action&#x3D;ModifyIssueComment | 事项评论修改
[**ModifyIssueDescription**](DefaultApi.md#ModifyIssueDescription) | **Post** /?action&#x3D;ModifyIssueDescription | 事项描述修改
[**ModifyIssueParentRequirement**](DefaultApi.md#ModifyIssueParentRequirement) | **Post** /?action&#x3D;ModifyIssueParentRequirement | 事项父需求修改
[**ModifyIteration**](DefaultApi.md#ModifyIteration) | **Post** /?action&#x3D;ModifyIteration | 迭代修改
[**ModifyMergeMR**](DefaultApi.md#ModifyMergeMR) | **Post** /ModifyMergeMR | 合并信息-执行合并
[**ModifyMergeRequestBasicSettings**](DefaultApi.md#ModifyMergeRequestBasicSettings) | **Post** /ModifyMergeRequestBasicSettings | 仓库设置-修改合并请求基础设置
[**ModifyMergeRequestMergeCommitMessageTemplate**](DefaultApi.md#ModifyMergeRequestMergeCommitMessageTemplate) | **Post** /ModifyMergeRequestMergeCommitMessageTemplate | 仓库设置-修改合并请求合并提交消息模版
[**ModifyMergeRequestSquashCommitMessageTemplate**](DefaultApi.md#ModifyMergeRequestSquashCommitMessageTemplate) | **Post** /ModifyMergeRequestSquashCommitMessageTemplate | 仓库设置-修改合并请求合并压缩提交消息模版
[**ModifyPolicy**](DefaultApi.md#ModifyPolicy) | **Post** /?action&#x3D;ModifyPolicy | 权限组修改
[**ModifyProject**](DefaultApi.md#ModifyProject) | **Post** /?action&#x3D;ModifyProject | 项目信息修改
[**ModifyProjectAnnouncement**](DefaultApi.md#ModifyProjectAnnouncement) | **Post** /?action&#x3D;ModifyProjectAnnouncement | 项目公告更新
[**ModifyProjectLabel**](DefaultApi.md#ModifyProjectLabel) | **Post** /?action&#x3D;ModifyProjectLabel | 项目标签修改
[**ModifyProjectPermission**](DefaultApi.md#ModifyProjectPermission) | **Post** /?action&#x3D;ModifyProjectPermission | 项目成员权限配置(老权限)
[**ModifyRelease**](DefaultApi.md#ModifyRelease) | **Post** /ModifyRelease | 版本修改
[**ModifyTeamLevelDepotSpec**](DefaultApi.md#ModifyTeamLevelDepotSpec) | **Post** /ModifyTeamLevelDepotSpec | 仓库设置-修改、新增团队级别的仓库规范
[**ModifyTeamMemberLocked**](DefaultApi.md#ModifyTeamMemberLocked) | **Post** /?action&#x3D;ModifyTeamMemberLocked | 团队成员锁定
[**ModifyTeamMemberUnlocked**](DefaultApi.md#ModifyTeamMemberUnlocked) | **Post** /?action&#x3D;ModifyTeamMemberUnlocked | 团队成员解锁
[**ModifyTestCase**](DefaultApi.md#ModifyTestCase) | **Post** /?action&#x3D;ModifyTestCase | 测试用例修改
[**ModifyTestCaseSection**](DefaultApi.md#ModifyTestCaseSection) | **Post** /?action&#x3D;ModifyTestCaseSection | 测试用例分组修改
[**ModifyTestRun**](DefaultApi.md#ModifyTestRun) | **Post** /?action&#x3D;ModifyTestRun | 测试计划修改
[**ModifyWorkItemSplitIssues**](DefaultApi.md#ModifyWorkItemSplitIssues) | **Post** /ModifyWorkItemSplitIssues | 项目集工作项分解&amp;取消分解到项目中的事项
[**PlanIterationIssue**](DefaultApi.md#PlanIterationIssue) | **Post** /?action&#x3D;PlanIterationIssue | 迭代批量规划
[**ReleaseArtifactVersion**](DefaultApi.md#ReleaseArtifactVersion) | **Post** /?action&#x3D;ReleaseArtifactVersion | 制品版本发布
[**ReorderCdPipelines**](DefaultApi.md#ReorderCdPipelines) | **Post** /?action&#x3D;ReorderCdPipelines | 部署流程重排序
[**SetGrantToResource**](DefaultApi.md#SetGrantToResource) | **Post** /?action&#x3D;SetGrantToResource | 授权设置，收回授权体在资源中的所有授权，重新设置为参数指定的授权
[**SetPredicatePolicyOnResource**](DefaultApi.md#SetPredicatePolicyOnResource) | **Post** /?action&#x3D;SetPredicatePolicyOnResource | 资源权限判定策略设置
[**StopCodingCIBuild**](DefaultApi.md#StopCodingCIBuild) | **Post** /?action&#x3D;StopCodingCIBuild | 构建停止
[**TriggerCdPipeline**](DefaultApi.md#TriggerCdPipeline) | **Post** /?action&#x3D;TriggerCdPipeline | 部署流程触发
[**TriggerCodingCIBuild**](DefaultApi.md#TriggerCodingCIBuild) | **Post** /?action&#x3D;TriggerCodingCIBuild | 构建触发
[**TriggerCodingScan**](DefaultApi.md#TriggerCodingScan) | **Post** /?action&#x3D;TriggerCodingScan | 代码扫描触发
[**UpdateUserGroupById**](DefaultApi.md#UpdateUserGroupById) | **Post** /?action&#x3D;UpdateUserGroupById | 用户组更新



## ArchiveTestRun

> ArchiveTestRun200Response ArchiveTestRun(ctx).Authorization(authorization).Action(action).ArchiveTestRunRequest(archiveTestRunRequest).Execute()

测试计划归档



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ArchiveTestRun" // string | Action (default to "ArchiveTestRun")
	archiveTestRunRequest := *openapiclient.NewArchiveTestRunRequest("ProjectName_example", int32(123)) // ArchiveTestRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ArchiveTestRun(context.Background()).Authorization(authorization).Action(action).ArchiveTestRunRequest(archiveTestRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ArchiveTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveTestRun`: ArchiveTestRun200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ArchiveTestRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ArchiveTestRun&quot;]
 **archiveTestRunRequest** | [**ArchiveTestRunRequest**](ArchiveTestRunRequest.md) |  | 

### Return type

[**ArchiveTestRun200Response**](ArchiveTestRun200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachResourceScopeToPolicy

> DeleteAPIDoc200Response AttachResourceScopeToPolicy(ctx).Authorization(authorization).Action(action).AttachResourceScopeToPolicyRequest(attachResourceScopeToPolicyRequest).Execute()

权限组添加可用的资源，原有其它资源不受影响，若已存在的资源不再进行追加



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "AttachResourceScopeToPolicy" // string | Action (default to "AttachResourceScopeToPolicy")
	attachResourceScopeToPolicyRequest := *openapiclient.NewAttachResourceScopeToPolicyRequest(int64(123), []openapiclient.ResourceInfoOfPolicyScope{*openapiclient.NewResourceInfoOfPolicyScope("ResourceId_example", "ResourceType_example")}) // AttachResourceScopeToPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.AttachResourceScopeToPolicy(context.Background()).Authorization(authorization).Action(action).AttachResourceScopeToPolicyRequest(attachResourceScopeToPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AttachResourceScopeToPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachResourceScopeToPolicy`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AttachResourceScopeToPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachResourceScopeToPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;AttachResourceScopeToPolicy&quot;]
 **attachResourceScopeToPolicyRequest** | [**AttachResourceScopeToPolicyRequest**](AttachResourceScopeToPolicyRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachToResource

> DeleteAPIDoc200Response AttachToResource(ctx).Authorization(authorization).Action(action).DetachFromResourceRequest(detachFromResourceRequest).Execute()

授权追加，原有其它授权不受影响，若授权已存在不再进行追加



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "AttachToResource" // string | Action (default to "AttachToResource")
	detachFromResourceRequest := *openapiclient.NewDetachFromResourceRequest([]openapiclient.GrantInfo{*openapiclient.NewGrantInfo("GrantObjectId_example", "GrantScope_example", int64(123))}, *openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example")) // DetachFromResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.AttachToResource(context.Background()).Authorization(authorization).Action(action).DetachFromResourceRequest(detachFromResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AttachToResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachToResource`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AttachToResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachToResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;AttachToResource&quot;]
 **detachFromResourceRequest** | [**DetachFromResourceRequest**](DetachFromResourceRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BindCdApplicationToProject

> DeleteAPIDoc200Response BindCdApplicationToProject(ctx).Authorization(authorization).Action(action).BindCdApplicationToProjectRequest(bindCdApplicationToProjectRequest).Execute()

绑定 CD 应用到项目



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "BindCdApplicationToProject" // string | Action (default to "BindCdApplicationToProject")
	bindCdApplicationToProjectRequest := *openapiclient.NewBindCdApplicationToProjectRequest("Application_example", "ProjectName_example") // BindCdApplicationToProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.BindCdApplicationToProject(context.Background()).Authorization(authorization).Action(action).BindCdApplicationToProjectRequest(bindCdApplicationToProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BindCdApplicationToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindCdApplicationToProject`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BindCdApplicationToProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindCdApplicationToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;BindCdApplicationToProject&quot;]
 **bindCdApplicationToProjectRequest** | [**BindCdApplicationToProjectRequest**](BindCdApplicationToProjectRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoundExternalDepot

> DeleteAPIDoc200Response BoundExternalDepot(ctx).Authorization(authorization).Action(action).BoundExternalDepotRequest(boundExternalDepotRequest).Execute()

外部仓库关联



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "BoundExternalDepot" // string | Action (default to "BoundExternalDepot")
	boundExternalDepotRequest := *openapiclient.NewBoundExternalDepotRequest("DepotType_example", "ExternalDepotAddress_example", int32(123), false) // BoundExternalDepotRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.BoundExternalDepot(context.Background()).Authorization(authorization).Action(action).BoundExternalDepotRequest(boundExternalDepotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BoundExternalDepot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoundExternalDepot`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BoundExternalDepot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBoundExternalDepotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;BoundExternalDepot&quot;]
 **boundExternalDepotRequest** | [**BoundExternalDepotRequest**](BoundExternalDepotRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelCdPipeline

> DeleteAPIDoc200Response CancelCdPipeline(ctx).Authorization(authorization).Action(action).CancelCdPipelineRequest(cancelCdPipelineRequest).Execute()

CD 部署流程取消执行



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CancelCdPipeline" // string | Action (default to "CancelCdPipeline")
	cancelCdPipelineRequest := *openapiclient.NewCancelCdPipelineRequest("PipelineExecutionId_example") // CancelCdPipelineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CancelCdPipeline(context.Background()).Authorization(authorization).Action(action).CancelCdPipelineRequest(cancelCdPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CancelCdPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelCdPipeline`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CancelCdPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelCdPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CancelCdPipeline&quot;]
 **cancelCdPipelineRequest** | [**CancelCdPipelineRequest**](CancelCdPipelineRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearCodingCIJobCache

> DeleteAPIDoc200Response ClearCodingCIJobCache(ctx).Authorization(authorization).Action(action).ClearCodingCIJobCacheRequest(clearCodingCIJobCacheRequest).Execute()

构建计划缓存清理



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ClearCodingCIJobCache" // string | Action (default to "ClearCodingCIJobCache")
	clearCodingCIJobCacheRequest := *openapiclient.NewClearCodingCIJobCacheRequest(int64(123)) // ClearCodingCIJobCacheRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ClearCodingCIJobCache(context.Background()).Authorization(authorization).Action(action).ClearCodingCIJobCacheRequest(clearCodingCIJobCacheRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClearCodingCIJobCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearCodingCIJobCache`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ClearCodingCIJobCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearCodingCIJobCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ClearCodingCIJobCache&quot;]
 **clearCodingCIJobCacheRequest** | [**ClearCodingCIJobCacheRequest**](ClearCodingCIJobCacheRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateArtifactCredit

> CreateArtifactCredit200Response CreateArtifactCredit(ctx).Authorization(authorization).Action(action).CreateArtifactCreditRequest(createArtifactCreditRequest).Execute()

制品授信清单创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateArtifactCredit" // string | Action (default to "CreateArtifactCredit")
	createArtifactCreditRequest := *openapiclient.NewCreateArtifactCreditRequest([]openapiclient.ArtifactsOpenApiCreateArtifactCreditsRangeData{*openapiclient.NewArtifactsOpenApiCreateArtifactCreditsRangeData("RangeType_example", int64(123))}, false, []openapiclient.ArtifactsOpenApiArtifactCreditsRuleData{*openapiclient.NewArtifactsOpenApiArtifactCreditsRuleData("Version_example", int32(123), "PkgNameAlgorithm_example", "PkgName_example", "VersionAlgorithm_example")}, "Name_example") // CreateArtifactCreditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateArtifactCredit(context.Background()).Authorization(authorization).Action(action).CreateArtifactCreditRequest(createArtifactCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateArtifactCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArtifactCredit`: CreateArtifactCredit200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateArtifactCredit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateArtifactCredit&quot;]
 **createArtifactCreditRequest** | [**CreateArtifactCreditRequest**](CreateArtifactCreditRequest.md) |  | 

### Return type

[**CreateArtifactCredit200Response**](CreateArtifactCredit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateArtifactProperties

> DeleteAPIDoc200Response CreateArtifactProperties(ctx).Authorization(authorization).Action(action).ModifyArtifactPropertiesRequest(modifyArtifactPropertiesRequest).Execute()

制品属性新增（指定版本）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateArtifactProperties" // string | Action (default to "CreateArtifactProperties")
	modifyArtifactPropertiesRequest := *openapiclient.NewModifyArtifactPropertiesRequest("Package_example", "PackageVersion_example", int32(123), []openapiclient.ArtifactPropertyBean{*openapiclient.NewArtifactPropertyBean("Name_example", "Value_example")}, "Repository_example") // ModifyArtifactPropertiesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateArtifactProperties(context.Background()).Authorization(authorization).Action(action).ModifyArtifactPropertiesRequest(modifyArtifactPropertiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateArtifactProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArtifactProperties`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateArtifactProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateArtifactProperties&quot;]
 **modifyArtifactPropertiesRequest** | [**ModifyArtifactPropertiesRequest**](ModifyArtifactPropertiesRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateArtifactRepository

> CreateArtifactRepository200Response CreateArtifactRepository(ctx).Authorization(authorization).Action(action).CreateArtifactRepositoryRequest(createArtifactRepositoryRequest).Execute()

制品仓库创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateArtifactRepository" // string | Action (default to "CreateArtifactRepository")
	createArtifactRepositoryRequest := *openapiclient.NewCreateArtifactRepositoryRequest(int32(123), "RepositoryName_example", int32(123)) // CreateArtifactRepositoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateArtifactRepository(context.Background()).Authorization(authorization).Action(action).CreateArtifactRepositoryRequest(createArtifactRepositoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateArtifactRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArtifactRepository`: CreateArtifactRepository200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateArtifactRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateArtifactRepository&quot;]
 **createArtifactRepositoryRequest** | [**CreateArtifactRepositoryRequest**](CreateArtifactRepositoryRequest.md) |  | 

### Return type

[**CreateArtifactRepository200Response**](CreateArtifactRepository200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttachmentPrepareSignUrl

> CreateAttachmentPrepareSignUrl200Response CreateAttachmentPrepareSignUrl(ctx).Authorization(authorization).Action(action).CreateAttachmentPrepareSignUrlRequest(createAttachmentPrepareSignUrlRequest).Execute()

附件预上传信息生成



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateAttachmentPrepareSignUrl" // string | Action (default to "CreateAttachmentPrepareSignUrl")
	createAttachmentPrepareSignUrlRequest := *openapiclient.NewCreateAttachmentPrepareSignUrlRequest("FileName_example", "ProjectName_example") // CreateAttachmentPrepareSignUrlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateAttachmentPrepareSignUrl(context.Background()).Authorization(authorization).Action(action).CreateAttachmentPrepareSignUrlRequest(createAttachmentPrepareSignUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAttachmentPrepareSignUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAttachmentPrepareSignUrl`: CreateAttachmentPrepareSignUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAttachmentPrepareSignUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentPrepareSignUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateAttachmentPrepareSignUrl&quot;]
 **createAttachmentPrepareSignUrlRequest** | [**CreateAttachmentPrepareSignUrlRequest**](CreateAttachmentPrepareSignUrlRequest.md) |  | 

### Return type

[**CreateAttachmentPrepareSignUrl200Response**](CreateAttachmentPrepareSignUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBinaryFiles

> CreateBinaryFiles200Response CreateBinaryFiles(ctx).Authorization(authorization).Action(action).CreateBinaryFilesRequest(createBinaryFilesRequest).Execute()

Git文件-Git仓库创建二进制文件



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateBinaryFiles" // string | Action (default to "CreateBinaryFiles")
	createBinaryFilesRequest := *openapiclient.NewCreateBinaryFilesRequest(int64(123), []openapiclient.GitFile{*openapiclient.NewGitFile("Content_example", "Path_example")}, "LastCommitSha_example", "Message_example", "SrcRef_example") // CreateBinaryFilesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateBinaryFiles(context.Background()).Authorization(authorization).Action(action).CreateBinaryFilesRequest(createBinaryFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBinaryFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBinaryFiles`: CreateBinaryFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBinaryFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBinaryFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateBinaryFiles&quot;]
 **createBinaryFilesRequest** | [**CreateBinaryFilesRequest**](CreateBinaryFilesRequest.md) |  | 

### Return type

[**CreateBinaryFiles200Response**](CreateBinaryFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBranchProtection

> DeleteAPIDoc200Response CreateBranchProtection(ctx).Authorization(authorization).Action(action).CreateBranchProtectionRequest(createBranchProtectionRequest).Execute()

仓库设置-新增代码保护规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateBranchProtection" // string | Action (default to "CreateBranchProtection")
	createBranchProtectionRequest := *openapiclient.NewCreateBranchProtectionRequest(int64(123), "Rule_example") // CreateBranchProtectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateBranchProtection(context.Background()).Authorization(authorization).Action(action).CreateBranchProtectionRequest(createBranchProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBranchProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBranchProtection`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBranchProtection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateBranchProtection&quot;]
 **createBranchProtectionRequest** | [**CreateBranchProtectionRequest**](CreateBranchProtectionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBranchProtectionMember

> DeleteAPIDoc200Response CreateBranchProtectionMember(ctx).Authorization(authorization).Action(action).CreateBranchProtectionMemberRequest(createBranchProtectionMemberRequest).Execute()

仓库设置-新增保护分支规则管理员



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateBranchProtectionMember" // string | Action (default to "CreateBranchProtectionMember")
	createBranchProtectionMemberRequest := *openapiclient.NewCreateBranchProtectionMemberRequest(false, int64(123), int64(123), "UserGlobalKey_example") // CreateBranchProtectionMemberRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateBranchProtectionMember(context.Background()).Authorization(authorization).Action(action).CreateBranchProtectionMemberRequest(createBranchProtectionMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBranchProtectionMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBranchProtectionMember`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBranchProtectionMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBranchProtectionMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateBranchProtectionMember&quot;]
 **createBranchProtectionMemberRequest** | [**CreateBranchProtectionMemberRequest**](CreateBranchProtectionMemberRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCaseResult

> DeleteAPIDoc200Response CreateCaseResult(ctx).Authorization(authorization).Action(action).CreateCaseResultRequest(createCaseResultRequest).Execute()

测试用例添加测试结果



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCaseResult" // string | Action (default to "CreateCaseResult")
	createCaseResultRequest := *openapiclient.NewCreateCaseResultRequest(int32(123), "ProjectName_example", int32(123), "Status_example") // CreateCaseResultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCaseResult(context.Background()).Authorization(authorization).Action(action).CreateCaseResultRequest(createCaseResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCaseResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCaseResult`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCaseResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCaseResult&quot;]
 **createCaseResultRequest** | [**CreateCaseResultRequest**](CreateCaseResultRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCdCloudAccount

> CreateCdCloudAccount200Response CreateCdCloudAccount(ctx).Authorization(authorization).Action(action).CreateCdCloudAccountRequest(createCdCloudAccountRequest).Execute()

CD 云账号添加



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCdCloudAccount" // string | Action (default to "CreateCdCloudAccount")
	createCdCloudAccountRequest := *openapiclient.NewCreateCdCloudAccountRequest("CloudProvider_example", *openapiclient.NewCloudAccountCredential(), "Name_example") // CreateCdCloudAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCdCloudAccount(context.Background()).Authorization(authorization).Action(action).CreateCdCloudAccountRequest(createCdCloudAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCdCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCdCloudAccount`: CreateCdCloudAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCdCloudAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCdCloudAccount&quot;]
 **createCdCloudAccountRequest** | [**CreateCdCloudAccountRequest**](CreateCdCloudAccountRequest.md) |  | 

### Return type

[**CreateCdCloudAccount200Response**](CreateCdCloudAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCdHostServerGroup

> CreateCdHostServerGroup200Response CreateCdHostServerGroup(ctx).Authorization(authorization).Action(action).CreateCdHostServerGroupRequest(createCdHostServerGroupRequest).Execute()

CD 主机组添加



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCdHostServerGroup" // string | Action (default to "CreateCdHostServerGroup")
	createCdHostServerGroupRequest := *openapiclient.NewCreateCdHostServerGroupRequest(int64(123), "AuthMethod_example", "DisplayName_example", []string{"Ips_example"}, int64(123), "UserName_example") // CreateCdHostServerGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCdHostServerGroup(context.Background()).Authorization(authorization).Action(action).CreateCdHostServerGroupRequest(createCdHostServerGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCdHostServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCdHostServerGroup`: CreateCdHostServerGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCdHostServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdHostServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCdHostServerGroup&quot;]
 **createCdHostServerGroupRequest** | [**CreateCdHostServerGroupRequest**](CreateCdHostServerGroupRequest.md) |  | 

### Return type

[**CreateCdHostServerGroup200Response**](CreateCdHostServerGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCdPipeline

> CreateCdPipeline200Response CreateCdPipeline(ctx).Authorization(authorization).Action(action).CreateCdPipelineRequest(createCdPipelineRequest).Execute()

CD 部署流程创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCdPipeline" // string | Action (default to "CreateCdPipeline")
	createCdPipelineRequest := *openapiclient.NewCreateCdPipelineRequest("PipelineJsonContent_example") // CreateCdPipelineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCdPipeline(context.Background()).Authorization(authorization).Action(action).CreateCdPipelineRequest(createCdPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCdPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCdPipeline`: CreateCdPipeline200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCdPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCdPipeline&quot;]
 **createCdPipelineRequest** | [**CreateCdPipelineRequest**](CreateCdPipelineRequest.md) |  | 

### Return type

[**CreateCdPipeline200Response**](CreateCdPipeline200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCdTask

> CreateCdTask200Response CreateCdTask(ctx).Authorization(authorization).Action(action).CreateCdTaskRequest(createCdTaskRequest).Execute()

CD 任务执行



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCdTask" // string | Action (default to "CreateCdTask")
	createCdTaskRequest := *openapiclient.NewCreateCdTaskRequest("TaskJsonContent_example") // CreateCdTaskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCdTask(context.Background()).Authorization(authorization).Action(action).CreateCdTaskRequest(createCdTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCdTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCdTask`: CreateCdTask200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCdTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCdTask&quot;]
 **createCdTaskRequest** | [**CreateCdTaskRequest**](CreateCdTaskRequest.md) |  | 

### Return type

[**CreateCdTask200Response**](CreateCdTask200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCodingCIJob

> CreateCodingCIJob200Response CreateCodingCIJob(ctx).Authorization(authorization).Action(action).CreateCodingCIJobRequest(createCodingCIJobRequest).Execute()

构建计划创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCodingCIJob" // string | Action (default to "CreateCodingCIJob")
	createCodingCIJobRequest := *openapiclient.NewCreateCodingCIJobRequest(false, false, int32(123), "DepotType_example", "ExecuteIn_example", "HookType_example", "JenkinsFileFromType_example", "JobFromType_example", "Name_example", int32(123)) // CreateCodingCIJobRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCodingCIJob(context.Background()).Authorization(authorization).Action(action).CreateCodingCIJobRequest(createCodingCIJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCodingCIJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCodingCIJob`: CreateCodingCIJob200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCodingCIJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCodingCIJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCodingCIJob&quot;]
 **createCodingCIJobRequest** | [**CreateCodingCIJobRequest**](CreateCodingCIJobRequest.md) |  | 

### Return type

[**CreateCodingCIJob200Response**](CreateCodingCIJob200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCodingCIJobByTeamTemplate

> CreateCodingCIJobByTeamTemplate200Response CreateCodingCIJobByTeamTemplate(ctx).Authorization(authorization).Action(action).CreateCodingCIJobByTeamTemplateRequest(createCodingCIJobByTeamTemplateRequest).Execute()

构建计划-根据团队模版创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCodingCIJobByTeamTemplate" // string | Action (default to "CreateCodingCIJobByTeamTemplate")
	createCodingCIJobByTeamTemplateRequest := *openapiclient.NewCreateCodingCIJobByTeamTemplateRequest(int64(123), "DepotType_example", "JobName_example", int64(123), int64(123)) // CreateCodingCIJobByTeamTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCodingCIJobByTeamTemplate(context.Background()).Authorization(authorization).Action(action).CreateCodingCIJobByTeamTemplateRequest(createCodingCIJobByTeamTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCodingCIJobByTeamTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCodingCIJobByTeamTemplate`: CreateCodingCIJobByTeamTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCodingCIJobByTeamTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCodingCIJobByTeamTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCodingCIJobByTeamTemplate&quot;]
 **createCodingCIJobByTeamTemplateRequest** | [**CreateCodingCIJobByTeamTemplateRequest**](CreateCodingCIJobByTeamTemplateRequest.md) |  | 

### Return type

[**CreateCodingCIJobByTeamTemplate200Response**](CreateCodingCIJobByTeamTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCodingProject

> CreateCodingProject200Response CreateCodingProject(ctx).Authorization(authorization).Action(action).CreateCodingProjectRequest(createCodingProjectRequest).Execute()

项目创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateCodingProject" // string | Action (default to "CreateCodingProject")
	createCodingProjectRequest := *openapiclient.NewCreateCodingProjectRequest("DisplayName_example", "Name_example", "ProjectTemplate_example", int32(123)) // CreateCodingProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateCodingProject(context.Background()).Authorization(authorization).Action(action).CreateCodingProjectRequest(createCodingProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCodingProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCodingProject`: CreateCodingProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCodingProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCodingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateCodingProject&quot;]
 **createCodingProjectRequest** | [**CreateCodingProjectRequest**](CreateCodingProjectRequest.md) |  | 

### Return type

[**CreateCodingProject200Response**](CreateCodingProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDepartment

> CreateDepartment200Response CreateDepartment(ctx).Authorization(authorization).Action(action).CreateDepartmentRequest(createDepartmentRequest).Execute()

部门创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateDepartment" // string | Action (default to "CreateDepartment")
	createDepartmentRequest := *openapiclient.NewCreateDepartmentRequest("Name_example", int64(123)) // CreateDepartmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateDepartment(context.Background()).Authorization(authorization).Action(action).CreateDepartmentRequest(createDepartmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDepartment`: CreateDepartment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDepartment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateDepartment&quot;]
 **createDepartmentRequest** | [**CreateDepartmentRequest**](CreateDepartmentRequest.md) |  | 

### Return type

[**CreateDepartment200Response**](CreateDepartment200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDepotByTemplate

> CreateDepotByTemplate200Response CreateDepotByTemplate(ctx).Authorization(authorization).Action(action).CreateDepotByTemplateRequest(createDepotByTemplateRequest).Execute()

仓库信息-根据模板创建仓库



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateDepotByTemplate" // string | Action (default to "CreateDepotByTemplate")
	createDepotByTemplateRequest := *openapiclient.NewCreateDepotByTemplateRequest("DepotName_example", int64(123), "Template_example", int64(123)) // CreateDepotByTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateDepotByTemplate(context.Background()).Authorization(authorization).Action(action).CreateDepotByTemplateRequest(createDepotByTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDepotByTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDepotByTemplate`: CreateDepotByTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDepotByTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepotByTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateDepotByTemplate&quot;]
 **createDepotByTemplateRequest** | [**CreateDepotByTemplateRequest**](CreateDepotByTemplateRequest.md) |  | 

### Return type

[**CreateDepotByTemplate200Response**](CreateDepotByTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDepotFilePushRule

> DeleteDepotFilePushRuleDenyPrivilege200Response CreateDepotFilePushRule(ctx).Authorization(authorization).Action(action).CreateDepotFilePushRuleRequest(createDepotFilePushRuleRequest).Execute()

仓库设置-新增git仓库文件推送规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateDepotFilePushRule" // string | Action (default to "CreateDepotFilePushRule")
	createDepotFilePushRuleRequest := *openapiclient.NewCreateDepotFilePushRuleRequest("DepotPath_example", false, "Pattern_example") // CreateDepotFilePushRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateDepotFilePushRule(context.Background()).Authorization(authorization).Action(action).CreateDepotFilePushRuleRequest(createDepotFilePushRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDepotFilePushRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDepotFilePushRule`: DeleteDepotFilePushRuleDenyPrivilege200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDepotFilePushRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepotFilePushRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateDepotFilePushRule&quot;]
 **createDepotFilePushRuleRequest** | [**CreateDepotFilePushRuleRequest**](CreateDepotFilePushRuleRequest.md) |  | 

### Return type

[**DeleteDepotFilePushRuleDenyPrivilege200Response**](DeleteDepotFilePushRuleDenyPrivilege200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDepotFilePushRulePrivilege

> DeleteDepotFilePushRuleDenyPrivilege200Response CreateDepotFilePushRulePrivilege(ctx).Authorization(authorization).Action(action).ModifyDepotFilePushRuleDenyPrivilegeRequest(modifyDepotFilePushRuleDenyPrivilegeRequest).Execute()

仓库设置-新增git仓库文件推送规则特权者



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateDepotFilePushRulePrivilege" // string | Action (default to "CreateDepotFilePushRulePrivilege")
	modifyDepotFilePushRuleDenyPrivilegeRequest := *openapiclient.NewModifyDepotFilePushRuleDenyPrivilegeRequest("DepotPath_example", int64(123), false, false, false, "UserGlobalKey_example") // ModifyDepotFilePushRuleDenyPrivilegeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateDepotFilePushRulePrivilege(context.Background()).Authorization(authorization).Action(action).ModifyDepotFilePushRuleDenyPrivilegeRequest(modifyDepotFilePushRuleDenyPrivilegeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDepotFilePushRulePrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDepotFilePushRulePrivilege`: DeleteDepotFilePushRuleDenyPrivilege200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDepotFilePushRulePrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepotFilePushRulePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateDepotFilePushRulePrivilege&quot;]
 **modifyDepotFilePushRuleDenyPrivilegeRequest** | [**ModifyDepotFilePushRuleDenyPrivilegeRequest**](ModifyDepotFilePushRuleDenyPrivilegeRequest.md) |  | 

### Return type

[**DeleteDepotFilePushRuleDenyPrivilege200Response**](DeleteDepotFilePushRuleDenyPrivilege200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFile

> CreateFile200Response CreateFile(ctx).Authorization(authorization).Action(action).CreateFileRequest(createFileRequest).Execute()

文件创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateFile" // string | Action (default to "CreateFile")
	createFileRequest := *openapiclient.NewCreateFileRequest("AuthToken_example", "StorageKey_example") // CreateFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateFile(context.Background()).Authorization(authorization).Action(action).CreateFileRequest(createFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFile`: CreateFile200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateFile&quot;]
 **createFileRequest** | [**CreateFileRequest**](CreateFileRequest.md) |  | 

### Return type

[**CreateFile200Response**](CreateFile200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolder

> CreateFolder200Response CreateFolder(ctx).Authorization(authorization).Action(action).CreateFolderRequest(createFolderRequest).Execute()

文件夹创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateFolder" // string | Action (default to "CreateFolder")
	createFolderRequest := *openapiclient.NewCreateFolderRequest("FoldName_example", int64(123), "ProjectName_example") // CreateFolderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateFolder(context.Background()).Authorization(authorization).Action(action).CreateFolderRequest(createFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: CreateFolder200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateFolder&quot;]
 **createFolderRequest** | [**CreateFolderRequest**](CreateFolderRequest.md) |  | 

### Return type

[**CreateFolder200Response**](CreateFolder200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitBranch

> DeleteAPIDoc200Response CreateGitBranch(ctx).Authorization(authorization).Action(action).CreateGitBranchRequest(createGitBranchRequest).Execute()

仓库分支-用于代码仓库新建分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitBranch" // string | Action (default to "CreateGitBranch")
	createGitBranchRequest := *openapiclient.NewCreateGitBranchRequest("BranchName_example", int64(123), "StartPoint_example") // CreateGitBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitBranch(context.Background()).Authorization(authorization).Action(action).CreateGitBranchRequest(createGitBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitBranch`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitBranch&quot;]
 **createGitBranchRequest** | [**CreateGitBranchRequest**](CreateGitBranchRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitCommit

> CreateGitCommit200Response CreateGitCommit(ctx).Authorization(authorization).Action(action).CreateGitCommitRequest(createGitCommitRequest).Execute()

Git提交-创建一次提交



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitCommit" // string | Action (default to "CreateGitCommit")
	createGitCommitRequest := *openapiclient.NewCreateGitCommitRequest([]openapiclient.CommitFile{*openapiclient.NewCommitFile("ContentEncoding_example", "Path_example", false, false)}, "DepotPath_example", "LastCommitSha_example", "Message_example", "Ref_example") // CreateGitCommitRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitCommit(context.Background()).Authorization(authorization).Action(action).CreateGitCommitRequest(createGitCommitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitCommit`: CreateGitCommit200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitCommit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitCommit&quot;]
 **createGitCommitRequest** | [**CreateGitCommitRequest**](CreateGitCommitRequest.md) |  | 

### Return type

[**CreateGitCommit200Response**](CreateGitCommit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitCommitComment

> CreateGitCommitComment200Response CreateGitCommitComment(ctx).Authorization(authorization).Action(action).CreateGitCommitCommentRequest(createGitCommitCommentRequest).Execute()

Git提交-为某次提交创建一条评论



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitCommitComment" // string | Action (default to "CreateGitCommitComment")
	createGitCommitCommentRequest := *openapiclient.NewCreateGitCommitCommentRequest("Content_example", int64(123), int64(123), "Path_example", "Sha_example") // CreateGitCommitCommentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitCommitComment(context.Background()).Authorization(authorization).Action(action).CreateGitCommitCommentRequest(createGitCommitCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitCommitComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitCommitComment`: CreateGitCommitComment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitCommitComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitCommitCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitCommitComment&quot;]
 **createGitCommitCommentRequest** | [**CreateGitCommitCommentRequest**](CreateGitCommitCommentRequest.md) |  | 

### Return type

[**CreateGitCommitComment200Response**](CreateGitCommitComment200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitCommitNote

> DeleteAPIDoc200Response CreateGitCommitNote(ctx).Authorization(authorization).Action(action).CreateGitCommitNoteRequest(createGitCommitNoteRequest).Execute()

Git提交-创建提交注释。注意：对于 NotesRef 参数建议默认为空，因为git会使用默认的ref ：refs/notes/commits，如果填了这个参数，会用这个参数指定的ref来保存您的git note，有可能会覆盖您原有的ref。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitCommitNote" // string | Action (default to "CreateGitCommitNote")
	createGitCommitNoteRequest := *openapiclient.NewCreateGitCommitNoteRequest("CommitMessage_example", "CommitSha_example", int64(123), "Note_example") // CreateGitCommitNoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitCommitNote(context.Background()).Authorization(authorization).Action(action).CreateGitCommitNoteRequest(createGitCommitNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitCommitNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitCommitNote`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitCommitNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitCommitNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitCommitNote&quot;]
 **createGitCommitNoteRequest** | [**CreateGitCommitNoteRequest**](CreateGitCommitNoteRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitDeployKey

> DeleteAPIDoc200Response CreateGitDeployKey(ctx).Authorization(authorization).Action(action).CreateGitDeployKeyRequest(createGitDeployKeyRequest).Execute()

仓库设置-新建部署公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitDeployKey" // string | Action (default to "CreateGitDeployKey")
	createGitDeployKeyRequest := *openapiclient.NewCreateGitDeployKeyRequest(false, "Content_example", int64(123), "Title_example") // CreateGitDeployKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitDeployKey(context.Background()).Authorization(authorization).Action(action).CreateGitDeployKeyRequest(createGitDeployKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitDeployKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitDeployKey`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitDeployKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitDeployKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitDeployKey&quot;]
 **createGitDeployKeyRequest** | [**CreateGitDeployKeyRequest**](CreateGitDeployKeyRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitDepot

> CreateGitDepot200Response CreateGitDepot(ctx).Authorization(authorization).Action(action).CreateGitDepotRequest(createGitDepotRequest).Execute()

仓库信息-创建代码仓库



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitDepot" // string | Action (default to "CreateGitDepot")
	createGitDepotRequest := *openapiclient.NewCreateGitDepotRequest("DepotName_example", int64(123)) // CreateGitDepotRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitDepot(context.Background()).Authorization(authorization).Action(action).CreateGitDepotRequest(createGitDepotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitDepot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitDepot`: CreateGitDepot200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitDepot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitDepotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitDepot&quot;]
 **createGitDepotRequest** | [**CreateGitDepotRequest**](CreateGitDepotRequest.md) |  | 

### Return type

[**CreateGitDepot200Response**](CreateGitDepot200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitFiles

> CreateGitFiles200Response CreateGitFiles(ctx).Authorization(authorization).Action(action).CreateGitFilesRequest(createGitFilesRequest).Execute()

Git文件-创建仓库文件



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitFiles" // string | Action (default to "CreateGitFiles")
	createGitFilesRequest := *openapiclient.NewCreateGitFilesRequest(int64(123), []openapiclient.GitFile{*openapiclient.NewGitFile("Content_example", "Path_example")}, "LastCommitSha_example", "Message_example", "Ref_example") // CreateGitFilesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitFiles(context.Background()).Authorization(authorization).Action(action).CreateGitFilesRequest(createGitFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitFiles`: CreateGitFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitFiles&quot;]
 **createGitFilesRequest** | [**CreateGitFilesRequest**](CreateGitFilesRequest.md) |  | 

### Return type

[**CreateGitFiles200Response**](CreateGitFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitMergeReq

> CreateGitMergeReq200Response CreateGitMergeReq(ctx).Authorization(authorization).Action(action).CreateGitMergeReqRequest(createGitMergeReqRequest).Execute()

合并请求-创建git合并请求



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitMergeReq" // string | Action (default to "CreateGitMergeReq")
	createGitMergeReqRequest := *openapiclient.NewCreateGitMergeReqRequest("Content_example", int64(123), "DestBranch_example", "SrcBranch_example", "Title_example") // CreateGitMergeReqRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitMergeReq(context.Background()).Authorization(authorization).Action(action).CreateGitMergeReqRequest(createGitMergeReqRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitMergeReq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitMergeReq`: CreateGitMergeReq200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitMergeReq`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitMergeReqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitMergeReq&quot;]
 **createGitMergeReqRequest** | [**CreateGitMergeReqRequest**](CreateGitMergeReqRequest.md) |  | 

### Return type

[**CreateGitMergeReq200Response**](CreateGitMergeReq200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitMergeRequest

> CreateGitMergeRequest200Response CreateGitMergeRequest(ctx).Authorization(authorization).Action(action).CreateGitMergeRequestRequest(createGitMergeRequestRequest).Execute()

合并请求-创建Git合并请求mr



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitMergeRequest" // string | Action (default to "CreateGitMergeRequest")
	createGitMergeRequestRequest := *openapiclient.NewCreateGitMergeRequestRequest("Content_example", int32(123), "DestBranch_example", "SrcBranch_example", "Title_example") // CreateGitMergeRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitMergeRequest(context.Background()).Authorization(authorization).Action(action).CreateGitMergeRequestRequest(createGitMergeRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitMergeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitMergeRequest`: CreateGitMergeRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitMergeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitMergeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitMergeRequest&quot;]
 **createGitMergeRequestRequest** | [**CreateGitMergeRequestRequest**](CreateGitMergeRequestRequest.md) |  | 

### Return type

[**CreateGitMergeRequest200Response**](CreateGitMergeRequest200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitProtectedTagRule

> DeleteAPIDoc200Response CreateGitProtectedTagRule(ctx).Authorization(authorization).Action(action).CreateGitProtectedTagRuleRequest(createGitProtectedTagRuleRequest).Execute()

仓库设置-创建标签保护规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitProtectedTagRule" // string | Action (default to "CreateGitProtectedTagRule")
	createGitProtectedTagRuleRequest := *openapiclient.NewCreateGitProtectedTagRuleRequest(int64(123), "Rule_example") // CreateGitProtectedTagRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitProtectedTagRule(context.Background()).Authorization(authorization).Action(action).CreateGitProtectedTagRuleRequest(createGitProtectedTagRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitProtectedTagRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitProtectedTagRule`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitProtectedTagRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitProtectedTagRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitProtectedTagRule&quot;]
 **createGitProtectedTagRuleRequest** | [**CreateGitProtectedTagRuleRequest**](CreateGitProtectedTagRuleRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitProtectedTagRules

> DeleteAPIDoc200Response CreateGitProtectedTagRules(ctx).Authorization(authorization).Action(action).CreateGitProtectedTagRulesRequest(createGitProtectedTagRulesRequest).Execute()

仓库设置-批量创建标签保护规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitProtectedTagRules" // string | Action (default to "CreateGitProtectedTagRules")
	createGitProtectedTagRulesRequest := *openapiclient.NewCreateGitProtectedTagRulesRequest(int64(123), []string{"Rule_example"}) // CreateGitProtectedTagRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitProtectedTagRules(context.Background()).Authorization(authorization).Action(action).CreateGitProtectedTagRulesRequest(createGitProtectedTagRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitProtectedTagRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitProtectedTagRules`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitProtectedTagRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitProtectedTagRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitProtectedTagRules&quot;]
 **createGitProtectedTagRulesRequest** | [**CreateGitProtectedTagRulesRequest**](CreateGitProtectedTagRulesRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitRelease

> DeleteAPIDoc200Response CreateGitRelease(ctx).Authorization(authorization).Action(action).CreateGitReleaseRequest(createGitReleaseRequest).Execute()

版本信息-新建git版本



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitRelease" // string | Action (default to "CreateGitRelease")
	createGitReleaseRequest := *openapiclient.NewCreateGitReleaseRequest("CommitSha_example", int64(123), "Description_example", false, "TagName_example", "TargetCommitBranch_example", "Title_example", int64(123)) // CreateGitReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitRelease(context.Background()).Authorization(authorization).Action(action).CreateGitReleaseRequest(createGitReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitRelease`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitRelease&quot;]
 **createGitReleaseRequest** | [**CreateGitReleaseRequest**](CreateGitReleaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitTag

> CreateGitTag200Response CreateGitTag(ctx).Authorization(authorization).Action(action).CreateGitTagRequest(createGitTagRequest).Execute()

标签信息-创建标签



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateGitTag" // string | Action (default to "CreateGitTag")
	createGitTagRequest := *openapiclient.NewCreateGitTagRequest(int64(123), "Message_example", "StartPoint_example", "TagName_example") // CreateGitTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateGitTag(context.Background()).Authorization(authorization).Action(action).CreateGitTagRequest(createGitTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGitTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitTag`: CreateGitTag200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGitTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateGitTag&quot;]
 **createGitTagRequest** | [**CreateGitTagRequest**](CreateGitTagRequest.md) |  | 

### Return type

[**CreateGitTag200Response**](CreateGitTag200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssue

> CreateIssue200Response CreateIssue(ctx).Authorization(authorization).Action(action).CreateIssueRequest(createIssueRequest).Execute()

事项创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateIssue" // string | Action (default to "CreateIssue")
	createIssueRequest := *openapiclient.NewCreateIssueRequest("Name_example", "Priority_example", "ProjectName_example", "Type_example") // CreateIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateIssue(context.Background()).Authorization(authorization).Action(action).CreateIssueRequest(createIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssue`: CreateIssue200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateIssue&quot;]
 **createIssueRequest** | [**CreateIssueRequest**](CreateIssueRequest.md) |  | 

### Return type

[**CreateIssue200Response**](CreateIssue200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueBlock

> DeleteAPIDoc200Response CreateIssueBlock(ctx).Authorization(authorization).Action(action).CreateIssueBlockRequest(createIssueBlockRequest).Execute()

前置事项添加



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateIssueBlock" // string | Action (default to "CreateIssueBlock")
	createIssueBlockRequest := *openapiclient.NewCreateIssueBlockRequest(int64(123), int64(123), "ProjectName_example") // CreateIssueBlockRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateIssueBlock(context.Background()).Authorization(authorization).Action(action).CreateIssueBlockRequest(createIssueBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIssueBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueBlock`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIssueBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateIssueBlock&quot;]
 **createIssueBlockRequest** | [**CreateIssueBlockRequest**](CreateIssueBlockRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueComment

> DeleteAPIDoc200Response CreateIssueComment(ctx).Authorization(authorization).Action(action).CreateIssueCommentRequest(createIssueCommentRequest).Execute()

事项评论创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateIssueComment" // string | Action (default to "CreateIssueComment")
	createIssueCommentRequest := *openapiclient.NewCreateIssueCommentRequest("Content_example", int64(123), int64(123), "ProjectName_example") // CreateIssueCommentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateIssueComment(context.Background()).Authorization(authorization).Action(action).CreateIssueCommentRequest(createIssueCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIssueComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueComment`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIssueComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateIssueComment&quot;]
 **createIssueCommentRequest** | [**CreateIssueCommentRequest**](CreateIssueCommentRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueModule

> CreateIssueModule200Response CreateIssueModule(ctx).Authorization(authorization).Action(action).CreateIssueModuleRequest(createIssueModuleRequest).Execute()

事项模块创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateIssueModule" // string | Action (default to "CreateIssueModule")
	createIssueModuleRequest := *openapiclient.NewCreateIssueModuleRequest("Name_example", "ProjectName_example") // CreateIssueModuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateIssueModule(context.Background()).Authorization(authorization).Action(action).CreateIssueModuleRequest(createIssueModuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIssueModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueModule`: CreateIssueModule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIssueModule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateIssueModule&quot;]
 **createIssueModuleRequest** | [**CreateIssueModuleRequest**](CreateIssueModuleRequest.md) |  | 

### Return type

[**CreateIssueModule200Response**](CreateIssueModule200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueWorkHours

> DeleteAPIDoc200Response CreateIssueWorkHours(ctx).Authorization(authorization).Action(action).CreateIssueWorkHoursRequest(createIssueWorkHoursRequest).Execute()

工时登记



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateIssueWorkHours" // string | Action (default to "CreateIssueWorkHours")
	createIssueWorkHoursRequest := *openapiclient.NewCreateIssueWorkHoursRequest(int64(123), "ProjectName_example", float32(123), float32(123), int64(123)) // CreateIssueWorkHoursRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateIssueWorkHours(context.Background()).Authorization(authorization).Action(action).CreateIssueWorkHoursRequest(createIssueWorkHoursRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIssueWorkHours``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueWorkHours`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIssueWorkHours`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueWorkHoursRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateIssueWorkHours&quot;]
 **createIssueWorkHoursRequest** | [**CreateIssueWorkHoursRequest**](CreateIssueWorkHoursRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIteration

> CreateIteration200Response CreateIteration(ctx).Authorization(authorization).Action(action).CreateIterationRequest(createIterationRequest).Execute()

迭代创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateIteration" // string | Action (default to "CreateIteration")
	createIterationRequest := *openapiclient.NewCreateIterationRequest("Name_example", "ProjectName_example") // CreateIterationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateIteration(context.Background()).Authorization(authorization).Action(action).CreateIterationRequest(createIterationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIteration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIteration`: CreateIteration200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIteration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateIteration&quot;]
 **createIterationRequest** | [**CreateIterationRequest**](CreateIterationRequest.md) |  | 

### Return type

[**CreateIteration200Response**](CreateIteration200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMemberSshKey

> CreateMemberSshKey200Response CreateMemberSshKey(ctx).Authorization(authorization).Action(action).CreateMemberSshKeyRequest(createMemberSshKeyRequest).Execute()

仓库设置-导入团队成员SSH公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateMemberSshKey" // string | Action (default to "CreateMemberSshKey")
	createMemberSshKeyRequest := *openapiclient.NewCreateMemberSshKeyRequest("Content_example", int64(123), "Title_example") // CreateMemberSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateMemberSshKey(context.Background()).Authorization(authorization).Action(action).CreateMemberSshKeyRequest(createMemberSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMemberSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMemberSshKey`: CreateMemberSshKey200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMemberSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateMemberSshKey&quot;]
 **createMemberSshKeyRequest** | [**CreateMemberSshKeyRequest**](CreateMemberSshKeyRequest.md) |  | 

### Return type

[**CreateMemberSshKey200Response**](CreateMemberSshKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMergeRequestNote

> CreateMergeRequestNote200Response CreateMergeRequestNote(ctx).Authorization(authorization).Action(action).CreateMergeRequestNoteRequest(createMergeRequestNoteRequest).Execute()

合并请求-创建合并请求行评论和改动文件diff行评论



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateMergeRequestNote" // string | Action (default to "CreateMergeRequestNote")
	createMergeRequestNoteRequest := *openapiclient.NewCreateMergeRequestNoteRequest("Content_example", "DepotPath_example", int32(123), int32(123)) // CreateMergeRequestNoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateMergeRequestNote(context.Background()).Authorization(authorization).Action(action).CreateMergeRequestNoteRequest(createMergeRequestNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMergeRequestNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMergeRequestNote`: CreateMergeRequestNote200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMergeRequestNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMergeRequestNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateMergeRequestNote&quot;]
 **createMergeRequestNoteRequest** | [**CreateMergeRequestNoteRequest**](CreateMergeRequestNoteRequest.md) |  | 

### Return type

[**CreateMergeRequestNote200Response**](CreateMergeRequestNote200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMergeRequestReviewer

> CreateMergeRequestReviewer200Response CreateMergeRequestReviewer(ctx).Authorization(authorization).Action(action).CreateMergeRequestReviewerRequest(createMergeRequestReviewerRequest).Execute()

合并请求-新增合并请求评审者



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateMergeRequestReviewer" // string | Action (default to "CreateMergeRequestReviewer")
	createMergeRequestReviewerRequest := *openapiclient.NewCreateMergeRequestReviewerRequest(int64(123), int64(123), "ReviewerGlobalKey_example") // CreateMergeRequestReviewerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateMergeRequestReviewer(context.Background()).Authorization(authorization).Action(action).CreateMergeRequestReviewerRequest(createMergeRequestReviewerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMergeRequestReviewer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMergeRequestReviewer`: CreateMergeRequestReviewer200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMergeRequestReviewer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMergeRequestReviewerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateMergeRequestReviewer&quot;]
 **createMergeRequestReviewerRequest** | [**CreateMergeRequestReviewerRequest**](CreateMergeRequestReviewerRequest.md) |  | 

### Return type

[**CreateMergeRequestReviewer200Response**](CreateMergeRequestReviewer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> ModifyPolicy200Response CreatePolicy(ctx).Authorization(authorization).Action(action).CreatePolicyRequest(createPolicyRequest).Execute()

权限组创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreatePolicy" // string | Action (default to "CreatePolicy")
	createPolicyRequest := *openapiclient.NewCreatePolicyRequest("Alias_example", "Name_example", *openapiclient.NewPolicyDocument([]openapiclient.PolicyStatement{*openapiclient.NewPolicyStatement([]string{"Action_example"}, "Effect_example", []string{"Resource_example"})}), "PolicyType_example", []string{"ResourceType_example"}) // CreatePolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreatePolicy(context.Background()).Authorization(authorization).Action(action).CreatePolicyRequest(createPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: ModifyPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreatePolicy&quot;]
 **createPolicyRequest** | [**CreatePolicyRequest**](CreatePolicyRequest.md) |  | 

### Return type

[**ModifyPolicy200Response**](ModifyPolicy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProgram

> CreateProgram200Response CreateProgram(ctx).Authorization(authorization).Action(action).CreateProgramRequest(createProgramRequest).Execute()

项目集创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProgram" // string | Action (default to "CreateProgram")
	createProgramRequest := *openapiclient.NewCreateProgramRequest("Description_example", "DisplayName_example", "2000-01-01", "Name_example", "2000-01-01") // CreateProgramRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProgram(context.Background()).Authorization(authorization).Action(action).CreateProgramRequest(createProgramRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProgram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProgram`: CreateProgram200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProgram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProgram&quot;]
 **createProgramRequest** | [**CreateProgramRequest**](CreateProgramRequest.md) |  | 

### Return type

[**CreateProgram200Response**](CreateProgram200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProgramMemberPolicy

> CreateProgramMemberPolicy200Response CreateProgramMemberPolicy(ctx).Authorization(authorization).Action(action).CreateProgramMemberPolicyRequest(createProgramMemberPolicyRequest).Execute()

项目集成员权限组添加



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProgramMemberPolicy" // string | Action (default to "CreateProgramMemberPolicy")
	createProgramMemberPolicyRequest := *openapiclient.NewCreateProgramMemberPolicyRequest([]string{"Policies_example"}, "19", "USER", int64(11)) // CreateProgramMemberPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProgramMemberPolicy(context.Background()).Authorization(authorization).Action(action).CreateProgramMemberPolicyRequest(createProgramMemberPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProgramMemberPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProgramMemberPolicy`: CreateProgramMemberPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProgramMemberPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramMemberPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProgramMemberPolicy&quot;]
 **createProgramMemberPolicyRequest** | [**CreateProgramMemberPolicyRequest**](CreateProgramMemberPolicyRequest.md) |  | 

### Return type

[**CreateProgramMemberPolicy200Response**](CreateProgramMemberPolicy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProgramProjects

> CreateProgramProjects200Response CreateProgramProjects(ctx).Authorization(authorization).Action(action).CreateProgramProjectsRequest(createProgramProjectsRequest).Execute()

项目集中添加项目



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProgramProjects" // string | Action (default to "CreateProgramProjects")
	createProgramProjectsRequest := *openapiclient.NewCreateProgramProjectsRequest(int64(123), []int64{int64(123)}) // CreateProgramProjectsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProgramProjects(context.Background()).Authorization(authorization).Action(action).CreateProgramProjectsRequest(createProgramProjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProgramProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProgramProjects`: CreateProgramProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProgramProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProgramProjects&quot;]
 **createProgramProjectsRequest** | [**CreateProgramProjectsRequest**](CreateProgramProjectsRequest.md) |  | 

### Return type

[**CreateProgramProjects200Response**](CreateProgramProjects200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectAnnouncement

> CreateProjectAnnouncement200Response CreateProjectAnnouncement(ctx).Authorization(authorization).Action(action).CreateProjectAnnouncementRequest(createProjectAnnouncementRequest).Execute()

项目公告创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProjectAnnouncement" // string | Action (default to "CreateProjectAnnouncement")
	createProjectAnnouncementRequest := *openapiclient.NewCreateProjectAnnouncementRequest("ProjectName_example", "Content_example") // CreateProjectAnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProjectAnnouncement(context.Background()).Authorization(authorization).Action(action).CreateProjectAnnouncementRequest(createProjectAnnouncementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProjectAnnouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectAnnouncement`: CreateProjectAnnouncement200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProjectAnnouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProjectAnnouncement&quot;]
 **createProjectAnnouncementRequest** | [**CreateProjectAnnouncementRequest**](CreateProjectAnnouncementRequest.md) |  | 

### Return type

[**CreateProjectAnnouncement200Response**](CreateProjectAnnouncement200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectLabel

> ModifyProjectLabel200Response CreateProjectLabel(ctx).Authorization(authorization).Action(action).CreateProjectLabelRequest(createProjectLabelRequest).Execute()

项目标签创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProjectLabel" // string | Action (default to "CreateProjectLabel")
	createProjectLabelRequest := *openapiclient.NewCreateProjectLabelRequest(int64(123), "Name_example", "Color_example") // CreateProjectLabelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProjectLabel(context.Background()).Authorization(authorization).Action(action).CreateProjectLabelRequest(createProjectLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProjectLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectLabel`: ModifyProjectLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProjectLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProjectLabel&quot;]
 **createProjectLabelRequest** | [**CreateProjectLabelRequest**](CreateProjectLabelRequest.md) |  | 

### Return type

[**ModifyProjectLabel200Response**](ModifyProjectLabel200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectMemberPrincipal

> DeleteAPIDoc200Response CreateProjectMemberPrincipal(ctx).Authorization(authorization).Action(action).CreateProjectMemberPrincipalRequest(createProjectMemberPrincipalRequest).Execute()

项目成员主体新增(包含用户组、部门、成员)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProjectMemberPrincipal" // string | Action (default to "CreateProjectMemberPrincipal")
	createProjectMemberPrincipalRequest := *openapiclient.NewCreateProjectMemberPrincipalRequest([]openapiclient.PrincipalAdd{*openapiclient.NewPrincipalAdd([]int64{int64(123)}, "PrincipalId_example", "PrincipalType_example")}, int32(123)) // CreateProjectMemberPrincipalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProjectMemberPrincipal(context.Background()).Authorization(authorization).Action(action).CreateProjectMemberPrincipalRequest(createProjectMemberPrincipalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProjectMemberPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectMemberPrincipal`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProjectMemberPrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectMemberPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProjectMemberPrincipal&quot;]
 **createProjectMemberPrincipalRequest** | [**CreateProjectMemberPrincipalRequest**](CreateProjectMemberPrincipalRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectWithTemplate

> CreateProjectWithTemplate200Response CreateProjectWithTemplate(ctx).Authorization(authorization).Action(action).CreateProjectWithTemplateRequest(createProjectWithTemplateRequest).Execute()

模版项目创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateProjectWithTemplate" // string | Action (default to "CreateProjectWithTemplate")
	createProjectWithTemplateRequest := *openapiclient.NewCreateProjectWithTemplateRequest("DisplayName_example", false, "Label_example", "Name_example", "ProjectTemplate_example", int32(123)) // CreateProjectWithTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateProjectWithTemplate(context.Background()).Authorization(authorization).Action(action).CreateProjectWithTemplateRequest(createProjectWithTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProjectWithTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectWithTemplate`: CreateProjectWithTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProjectWithTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectWithTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateProjectWithTemplate&quot;]
 **createProjectWithTemplateRequest** | [**CreateProjectWithTemplateRequest**](CreateProjectWithTemplateRequest.md) |  | 

### Return type

[**CreateProjectWithTemplate200Response**](CreateProjectWithTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReadOnlyRef

> DeleteAPIDoc200Response CreateReadOnlyRef(ctx).Authorization(authorization).Action(action).CreateReadOnlyRefRequest(createReadOnlyRefRequest).Execute()

仓库分支-创建只读分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateReadOnlyRef" // string | Action (default to "CreateReadOnlyRef")
	createReadOnlyRefRequest := *openapiclient.NewCreateReadOnlyRefRequest("DepotPath_example", "RefName_example") // CreateReadOnlyRefRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateReadOnlyRef(context.Background()).Authorization(authorization).Action(action).CreateReadOnlyRefRequest(createReadOnlyRefRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateReadOnlyRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReadOnlyRef`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateReadOnlyRef`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReadOnlyRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateReadOnlyRef&quot;]
 **createReadOnlyRefRequest** | [**CreateReadOnlyRefRequest**](CreateReadOnlyRefRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRelease

> ModifyRelease200Response CreateRelease(ctx).Authorization(authorization).Action(action).CreateReleaseRequest(createReleaseRequest).Execute()

版本创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateRelease" // string | Action (default to "CreateRelease")
	createReleaseRequest := *openapiclient.NewCreateReleaseRequest("Name_example", "ProjectName_example") // CreateReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateRelease(context.Background()).Authorization(authorization).Action(action).CreateReleaseRequest(createReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRelease`: ModifyRelease200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateRelease&quot;]
 **createReleaseRequest** | [**CreateReleaseRequest**](CreateReleaseRequest.md) |  | 

### Return type

[**ModifyRelease200Response**](ModifyRelease200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReport

> CreateReport200Response CreateReport(ctx).Authorization(authorization).Action(action).CreateReportRequest(createReportRequest).Execute()

测试报告创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateReport" // string | Action (default to "CreateReport")
	createReportRequest := *openapiclient.NewCreateReportRequest("Name_example", "ProjectName_example", []int32{int32(123)}) // CreateReportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateReport(context.Background()).Authorization(authorization).Action(action).CreateReportRequest(createReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReport`: CreateReport200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateReport&quot;]
 **createReportRequest** | [**CreateReportRequest**](CreateReportRequest.md) |  | 

### Return type

[**CreateReport200Response**](CreateReport200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequirementDefectRelation

> DeleteAPIDoc200Response CreateRequirementDefectRelation(ctx).Authorization(authorization).Action(action).CreateRequirementDefectRelationRequest(createRequirementDefectRelationRequest).Execute()

需求关联缺陷



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateRequirementDefectRelation" // string | Action (default to "CreateRequirementDefectRelation")
	createRequirementDefectRelationRequest := *openapiclient.NewCreateRequirementDefectRelationRequest(int64(123), "ProjectName_example", int64(123)) // CreateRequirementDefectRelationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateRequirementDefectRelation(context.Background()).Authorization(authorization).Action(action).CreateRequirementDefectRelationRequest(createRequirementDefectRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRequirementDefectRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequirementDefectRelation`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRequirementDefectRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequirementDefectRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateRequirementDefectRelation&quot;]
 **createRequirementDefectRelationRequest** | [**CreateRequirementDefectRelationRequest**](CreateRequirementDefectRelationRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSshKey

> CreateSshKey200Response CreateSshKey(ctx).Authorization(authorization).Action(action).CreateSshKeyRequest(createSshKeyRequest).Execute()

仓库设置-导入用户SSH公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateSshKey" // string | Action (default to "CreateSshKey")
	createSshKeyRequest := *openapiclient.NewCreateSshKeyRequest("Content_example", "Title_example") // CreateSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateSshKey(context.Background()).Authorization(authorization).Action(action).CreateSshKeyRequest(createSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSshKey`: CreateSshKey200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateSshKey&quot;]
 **createSshKeyRequest** | [**CreateSshKeyRequest**](CreateSshKeyRequest.md) |  | 

### Return type

[**CreateSshKey200Response**](CreateSshKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestCase

> CreateTestCase200Response CreateTestCase(ctx).Authorization(authorization).Action(action).CreateTestCaseRequest(createTestCaseRequest).Execute()

测试用例创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestCase" // string | Action (default to "CreateTestCase")
	createTestCaseRequest := *openapiclient.NewCreateTestCaseRequest("ProjectName_example", int32(123), "TemplateType_example", "Title_example") // CreateTestCaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestCase(context.Background()).Authorization(authorization).Action(action).CreateTestCaseRequest(createTestCaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestCase`: CreateTestCase200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestCase&quot;]
 **createTestCaseRequest** | [**CreateTestCaseRequest**](CreateTestCaseRequest.md) |  | 

### Return type

[**CreateTestCase200Response**](CreateTestCase200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestCaseSection

> ModifyTestCaseSection200Response CreateTestCaseSection(ctx).Authorization(authorization).Action(action).CreateTestCaseSectionRequest(createTestCaseSectionRequest).Execute()

测试用例分组创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestCaseSection" // string | Action (default to "CreateTestCaseSection")
	createTestCaseSectionRequest := *openapiclient.NewCreateTestCaseSectionRequest("Name_example", "ProjectName_example") // CreateTestCaseSectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestCaseSection(context.Background()).Authorization(authorization).Action(action).CreateTestCaseSectionRequest(createTestCaseSectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestCaseSection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestCaseSection`: ModifyTestCaseSection200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestCaseSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestCaseSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestCaseSection&quot;]
 **createTestCaseSectionRequest** | [**CreateTestCaseSectionRequest**](CreateTestCaseSectionRequest.md) |  | 

### Return type

[**ModifyTestCaseSection200Response**](ModifyTestCaseSection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestDefect

> DeleteAPIDoc200Response CreateTestDefect(ctx).Authorization(authorization).Action(action).CreateTestDefectRequest(createTestDefectRequest).Execute()

测试任务关联缺陷



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestDefect" // string | Action (default to "CreateTestDefect")
	createTestDefectRequest := *openapiclient.NewCreateTestDefectRequest(int32(123), "ProjectName_example", int32(123)) // CreateTestDefectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestDefect(context.Background()).Authorization(authorization).Action(action).CreateTestDefectRequest(createTestDefectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestDefect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestDefect`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestDefect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestDefectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestDefect&quot;]
 **createTestDefectRequest** | [**CreateTestDefectRequest**](CreateTestDefectRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestResult

> DeleteAPIDoc200Response CreateTestResult(ctx).Authorization(authorization).Action(action).CreateTestResultRequest(createTestResultRequest).Execute()

测试任务添加测试结果



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestResult" // string | Action (default to "CreateTestResult")
	createTestResultRequest := *openapiclient.NewCreateTestResultRequest("ProjectName_example", "Status_example", int32(123)) // CreateTestResultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestResult(context.Background()).Authorization(authorization).Action(action).CreateTestResultRequest(createTestResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestResult`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestResult&quot;]
 **createTestResultRequest** | [**CreateTestResultRequest**](CreateTestResultRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestResults

> DeleteAPIDoc200Response CreateTestResults(ctx).Authorization(authorization).Action(action).CreateTestResultsRequest(createTestResultsRequest).Execute()

测试任务状态批量更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestResults" // string | Action (default to "CreateTestResults")
	createTestResultsRequest := *openapiclient.NewCreateTestResultsRequest([]int32{int32(123)}, "ProjectName_example", int32(123), "Status_example") // CreateTestResultsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestResults(context.Background()).Authorization(authorization).Action(action).CreateTestResultsRequest(createTestResultsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestResults`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestResults&quot;]
 **createTestResultsRequest** | [**CreateTestResultsRequest**](CreateTestResultsRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestRun

> ModifyTestRun200Response CreateTestRun(ctx).Authorization(authorization).Action(action).CreateTestRunRequest(createTestRunRequest).Execute()

测试计划创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestRun" // string | Action (default to "CreateTestRun")
	createTestRunRequest := *openapiclient.NewCreateTestRunRequest(false, "Name_example", "ProjectName_example") // CreateTestRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestRun(context.Background()).Authorization(authorization).Action(action).CreateTestRunRequest(createTestRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestRun`: ModifyTestRun200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestRun&quot;]
 **createTestRunRequest** | [**CreateTestRunRequest**](CreateTestRunRequest.md) |  | 

### Return type

[**ModifyTestRun200Response**](ModifyTestRun200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTestStepResult

> DeleteAPIDoc200Response CreateTestStepResult(ctx).Authorization(authorization).Action(action).CreateTestStepResultRequest(createTestStepResultRequest).Execute()

测试任务添加某步骤的测试结果



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateTestStepResult" // string | Action (default to "CreateTestStepResult")
	createTestStepResultRequest := *openapiclient.NewCreateTestStepResultRequest("ProjectName_example", int32(123), "StepStatus_example", int32(123)) // CreateTestStepResultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateTestStepResult(context.Background()).Authorization(authorization).Action(action).CreateTestStepResultRequest(createTestStepResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTestStepResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTestStepResult`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTestStepResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestStepResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateTestStepResult&quot;]
 **createTestStepResultRequest** | [**CreateTestStepResultRequest**](CreateTestStepResultRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserGroup

> CreateUserGroup200Response CreateUserGroup(ctx).Authorization(authorization).Action(action).CreateUserGroupRequest(createUserGroupRequest).Execute()

用户组创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateUserGroup" // string | Action (default to "CreateUserGroup")
	createUserGroupRequest := *openapiclient.NewCreateUserGroupRequest("Description_example", "Name_example") // CreateUserGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateUserGroup(context.Background()).Authorization(authorization).Action(action).CreateUserGroupRequest(createUserGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserGroup`: CreateUserGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateUserGroup&quot;]
 **createUserGroupRequest** | [**CreateUserGroupRequest**](CreateUserGroupRequest.md) |  | 

### Return type

[**CreateUserGroup200Response**](CreateUserGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserGroupUsers

> DeleteAPIDoc200Response CreateUserGroupUsers(ctx).Authorization(authorization).Action(action).CreateUserGroupUsersRequest(createUserGroupUsersRequest).Execute()

用户组添加用户



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateUserGroupUsers" // string | Action (default to "CreateUserGroupUsers")
	createUserGroupUsersRequest := *openapiclient.NewCreateUserGroupUsersRequest([]openapiclient.UserGroupUserInfos{*openapiclient.NewUserGroupUserInfos(NullableInt64(123), NullableInt64(123))}) // CreateUserGroupUsersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateUserGroupUsers(context.Background()).Authorization(authorization).Action(action).CreateUserGroupUsersRequest(createUserGroupUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUserGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserGroupUsers`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUserGroupUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateUserGroupUsers&quot;]
 **createUserGroupUsersRequest** | [**CreateUserGroupUsersRequest**](CreateUserGroupUsersRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllUsersOnGroup

> DeleteAPIDoc200Response DeleteAllUsersOnGroup(ctx).Authorization(authorization).Action(action).DeleteAllUsersOnGroupRequest(deleteAllUsersOnGroupRequest).Execute()

用户组清理用户



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteAllUsersOnGroup" // string | Action (default to "DeleteAllUsersOnGroup")
	deleteAllUsersOnGroupRequest := *openapiclient.NewDeleteAllUsersOnGroupRequest(int64(123)) // DeleteAllUsersOnGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteAllUsersOnGroup(context.Background()).Authorization(authorization).Action(action).DeleteAllUsersOnGroupRequest(deleteAllUsersOnGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAllUsersOnGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllUsersOnGroup`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAllUsersOnGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllUsersOnGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteAllUsersOnGroup&quot;]
 **deleteAllUsersOnGroupRequest** | [**DeleteAllUsersOnGroupRequest**](DeleteAllUsersOnGroupRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactProperties

> DeleteAPIDoc200Response DeleteArtifactProperties(ctx).Authorization(authorization).Action(action).DeleteArtifactPropertiesRequest(deleteArtifactPropertiesRequest).Execute()

制品属性删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteArtifactProperties" // string | Action (default to "DeleteArtifactProperties")
	deleteArtifactPropertiesRequest := *openapiclient.NewDeleteArtifactPropertiesRequest("Package_example", "PackageVersion_example", int32(123), []string{"PropertyNameSet_example"}, "Repository_example") // DeleteArtifactPropertiesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteArtifactProperties(context.Background()).Authorization(authorization).Action(action).DeleteArtifactPropertiesRequest(deleteArtifactPropertiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteArtifactProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteArtifactProperties`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteArtifactProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteArtifactProperties&quot;]
 **deleteArtifactPropertiesRequest** | [**DeleteArtifactPropertiesRequest**](DeleteArtifactPropertiesRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBranchProtection

> DeleteAPIDoc200Response DeleteBranchProtection(ctx).Authorization(authorization).Action(action).DeleteBranchProtectionRequest(deleteBranchProtectionRequest).Execute()

仓库设置-删除保护分支规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteBranchProtection" // string | Action (default to "DeleteBranchProtection")
	deleteBranchProtectionRequest := *openapiclient.NewDeleteBranchProtectionRequest(int64(123), int64(123)) // DeleteBranchProtectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteBranchProtection(context.Background()).Authorization(authorization).Action(action).DeleteBranchProtectionRequest(deleteBranchProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBranchProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBranchProtection`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteBranchProtection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteBranchProtection&quot;]
 **deleteBranchProtectionRequest** | [**DeleteBranchProtectionRequest**](DeleteBranchProtectionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBranchProtectionMember

> DeleteAPIDoc200Response DeleteBranchProtectionMember(ctx).Authorization(authorization).Action(action).DeleteBranchProtectionMemberRequest(deleteBranchProtectionMemberRequest).Execute()

仓库设置-删除保护分支规则管理员



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteBranchProtectionMember" // string | Action (default to "DeleteBranchProtectionMember")
	deleteBranchProtectionMemberRequest := *openapiclient.NewDeleteBranchProtectionMemberRequest(int64(123), int64(123), "UserGlobalKey_example") // DeleteBranchProtectionMemberRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteBranchProtectionMember(context.Background()).Authorization(authorization).Action(action).DeleteBranchProtectionMemberRequest(deleteBranchProtectionMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBranchProtectionMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBranchProtectionMember`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteBranchProtectionMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBranchProtectionMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteBranchProtectionMember&quot;]
 **deleteBranchProtectionMemberRequest** | [**DeleteBranchProtectionMemberRequest**](DeleteBranchProtectionMemberRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdCloudAccount

> DeleteAPIDoc200Response DeleteCdCloudAccount(ctx).Authorization(authorization).Action(action).DeleteCdCloudAccountRequest(deleteCdCloudAccountRequest).Execute()

CD 云账号删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteCdCloudAccount" // string | Action (default to "DeleteCdCloudAccount")
	deleteCdCloudAccountRequest := *openapiclient.NewDeleteCdCloudAccountRequest(int64(123)) // DeleteCdCloudAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteCdCloudAccount(context.Background()).Authorization(authorization).Action(action).DeleteCdCloudAccountRequest(deleteCdCloudAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCdCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCdCloudAccount`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCdCloudAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteCdCloudAccount&quot;]
 **deleteCdCloudAccountRequest** | [**DeleteCdCloudAccountRequest**](DeleteCdCloudAccountRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdHostServerGroup

> DeleteAPIDoc200Response DeleteCdHostServerGroup(ctx).Authorization(authorization).Action(action).DeleteCdHostServerGroupRequest(deleteCdHostServerGroupRequest).Execute()

CD 主机组删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteCdHostServerGroup" // string | Action (default to "DeleteCdHostServerGroup")
	deleteCdHostServerGroupRequest := *openapiclient.NewDeleteCdHostServerGroupRequest(int64(123)) // DeleteCdHostServerGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteCdHostServerGroup(context.Background()).Authorization(authorization).Action(action).DeleteCdHostServerGroupRequest(deleteCdHostServerGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCdHostServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCdHostServerGroup`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCdHostServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdHostServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteCdHostServerGroup&quot;]
 **deleteCdHostServerGroupRequest** | [**DeleteCdHostServerGroupRequest**](DeleteCdHostServerGroupRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdPipeline

> DeleteAPIDoc200Response DeleteCdPipeline(ctx).Authorization(authorization).Action(action).DeleteCdPipelineRequest(deleteCdPipelineRequest).Execute()

CD 部署流程删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteCdPipeline" // string | Action (default to "DeleteCdPipeline")
	deleteCdPipelineRequest := *openapiclient.NewDeleteCdPipelineRequest("Application_example", "PipelineName_example") // DeleteCdPipelineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteCdPipeline(context.Background()).Authorization(authorization).Action(action).DeleteCdPipelineRequest(deleteCdPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCdPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCdPipeline`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCdPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteCdPipeline&quot;]
 **deleteCdPipelineRequest** | [**DeleteCdPipelineRequest**](DeleteCdPipelineRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCodingCIBuild

> DeleteAPIDoc200Response DeleteCodingCIBuild(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()

构建删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteCodingCIBuild" // string | Action (default to "DeleteCodingCIBuild")
	describeCodingCIBuildStageRequest := *openapiclient.NewDescribeCodingCIBuildStageRequest(int32(123)) // DescribeCodingCIBuildStageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteCodingCIBuild(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCodingCIBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCodingCIBuild`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCodingCIBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCodingCIBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteCodingCIBuild&quot;]
 **describeCodingCIBuildStageRequest** | [**DescribeCodingCIBuildStageRequest**](DescribeCodingCIBuildStageRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCodingCIJob

> DeleteAPIDoc200Response DeleteCodingCIJob(ctx).Authorization(authorization).Action(action).DeleteCodingCIJobRequest(deleteCodingCIJobRequest).Execute()

构建计划删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteCodingCIJob" // string | Action (default to "DeleteCodingCIJob")
	deleteCodingCIJobRequest := *openapiclient.NewDeleteCodingCIJobRequest(int32(123)) // DeleteCodingCIJobRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteCodingCIJob(context.Background()).Authorization(authorization).Action(action).DeleteCodingCIJobRequest(deleteCodingCIJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCodingCIJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCodingCIJob`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCodingCIJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCodingCIJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteCodingCIJob&quot;]
 **deleteCodingCIJobRequest** | [**DeleteCodingCIJobRequest**](DeleteCodingCIJobRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDepartment

> DeleteAPIDoc200Response DeleteDepartment(ctx).Authorization(authorization).Action(action).DeleteDepartmentRequest(deleteDepartmentRequest).Execute()

部门删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteDepartment" // string | Action (default to "DeleteDepartment")
	deleteDepartmentRequest := *openapiclient.NewDeleteDepartmentRequest(int64(123)) // DeleteDepartmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteDepartment(context.Background()).Authorization(authorization).Action(action).DeleteDepartmentRequest(deleteDepartmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDepartment`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteDepartment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteDepartment&quot;]
 **deleteDepartmentRequest** | [**DeleteDepartmentRequest**](DeleteDepartmentRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDepotFilePushRule

> DeleteDepotFilePushRule200Response DeleteDepotFilePushRule(ctx).Authorization(authorization).Action(action).DeleteDepotFilePushRuleRequest(deleteDepotFilePushRuleRequest).Execute()

仓库设置-删除git仓库文件推送规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteDepotFilePushRule" // string | Action (default to "DeleteDepotFilePushRule")
	deleteDepotFilePushRuleRequest := *openapiclient.NewDeleteDepotFilePushRuleRequest("DepotPath_example", int64(123)) // DeleteDepotFilePushRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteDepotFilePushRule(context.Background()).Authorization(authorization).Action(action).DeleteDepotFilePushRuleRequest(deleteDepotFilePushRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDepotFilePushRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDepotFilePushRule`: DeleteDepotFilePushRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteDepotFilePushRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDepotFilePushRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteDepotFilePushRule&quot;]
 **deleteDepotFilePushRuleRequest** | [**DeleteDepotFilePushRuleRequest**](DeleteDepotFilePushRuleRequest.md) |  | 

### Return type

[**DeleteDepotFilePushRule200Response**](DeleteDepotFilePushRule200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDepotFilePushRuleDenyPrivilege

> DeleteDepotFilePushRuleDenyPrivilege200Response DeleteDepotFilePushRuleDenyPrivilege(ctx).Authorization(authorization).Action(action).DeleteDepotFilePushRuleDenyPrivilegeRequest(deleteDepotFilePushRuleDenyPrivilegeRequest).Execute()

仓库设置-删除git仓库特权者文件推送权限



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteDepotFilePushRuleDenyPrivilege" // string | Action (default to "DeleteDepotFilePushRuleDenyPrivilege")
	deleteDepotFilePushRuleDenyPrivilegeRequest := *openapiclient.NewDeleteDepotFilePushRuleDenyPrivilegeRequest("DepotPath_example", int64(123), false, false, "UserGlobalKey_example") // DeleteDepotFilePushRuleDenyPrivilegeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteDepotFilePushRuleDenyPrivilege(context.Background()).Authorization(authorization).Action(action).DeleteDepotFilePushRuleDenyPrivilegeRequest(deleteDepotFilePushRuleDenyPrivilegeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDepotFilePushRuleDenyPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDepotFilePushRuleDenyPrivilege`: DeleteDepotFilePushRuleDenyPrivilege200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteDepotFilePushRuleDenyPrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDepotFilePushRuleDenyPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteDepotFilePushRuleDenyPrivilege&quot;]
 **deleteDepotFilePushRuleDenyPrivilegeRequest** | [**DeleteDepotFilePushRuleDenyPrivilegeRequest**](DeleteDepotFilePushRuleDenyPrivilegeRequest.md) |  | 

### Return type

[**DeleteDepotFilePushRuleDenyPrivilege200Response**](DeleteDepotFilePushRuleDenyPrivilege200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitBranch

> DeleteAPIDoc200Response DeleteGitBranch(ctx).Authorization(authorization).Action(action).DeleteGitBranchRequest(deleteGitBranchRequest).Execute()

仓库分支-删除代码仓库分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitBranch" // string | Action (default to "DeleteGitBranch")
	deleteGitBranchRequest := *openapiclient.NewDeleteGitBranchRequest("BranchName_example", int64(123)) // DeleteGitBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitBranch(context.Background()).Authorization(authorization).Action(action).DeleteGitBranchRequest(deleteGitBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitBranch`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitBranch&quot;]
 **deleteGitBranchRequest** | [**DeleteGitBranchRequest**](DeleteGitBranchRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitDeployKey

> DeleteAPIDoc200Response DeleteGitDeployKey(ctx).Authorization(authorization).Action(action).DeleteGitDeployKeyRequest(deleteGitDeployKeyRequest).Execute()

仓库设置-删除部署公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitDeployKey" // string | Action (default to "DeleteGitDeployKey")
	deleteGitDeployKeyRequest := *openapiclient.NewDeleteGitDeployKeyRequest(int64(123), int64(123)) // DeleteGitDeployKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitDeployKey(context.Background()).Authorization(authorization).Action(action).DeleteGitDeployKeyRequest(deleteGitDeployKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitDeployKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitDeployKey`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitDeployKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitDeployKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitDeployKey&quot;]
 **deleteGitDeployKeyRequest** | [**DeleteGitDeployKeyRequest**](DeleteGitDeployKeyRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitDepot

> DeleteAPIDoc200Response DeleteGitDepot(ctx).Authorization(authorization).Action(action).DeleteGitDepotRequest(deleteGitDepotRequest).Execute()

仓库信息-删除git仓库



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitDepot" // string | Action (default to "DeleteGitDepot")
	deleteGitDepotRequest := *openapiclient.NewDeleteGitDepotRequest(int64(123)) // DeleteGitDepotRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitDepot(context.Background()).Authorization(authorization).Action(action).DeleteGitDepotRequest(deleteGitDepotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitDepot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitDepot`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitDepot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitDepotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitDepot&quot;]
 **deleteGitDepotRequest** | [**DeleteGitDepotRequest**](DeleteGitDepotRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitFiles

> DeleteGitFiles200Response DeleteGitFiles(ctx).Authorization(authorization).Action(action).DeleteGitFilesRequest(deleteGitFilesRequest).Execute()

Git文件-删除文件并提交



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitFiles" // string | Action (default to "DeleteGitFiles")
	deleteGitFilesRequest := *openapiclient.NewDeleteGitFilesRequest("CommitMessage_example", int64(123), []string{"Paths_example"}, "Ref_example") // DeleteGitFilesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitFiles(context.Background()).Authorization(authorization).Action(action).DeleteGitFilesRequest(deleteGitFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitFiles`: DeleteGitFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitFiles&quot;]
 **deleteGitFilesRequest** | [**DeleteGitFilesRequest**](DeleteGitFilesRequest.md) |  | 

### Return type

[**DeleteGitFiles200Response**](DeleteGitFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitMergedBranches

> DeleteAPIDoc200Response DeleteGitMergedBranches(ctx).Authorization(authorization).Action(action).DeleteGitMergedBranchesRequest(deleteGitMergedBranchesRequest).Execute()

仓库分支-删除已合并到默认分支的分支（此操作不会删除受保护的分支)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitMergedBranches" // string | Action (default to "DeleteGitMergedBranches")
	deleteGitMergedBranchesRequest := *openapiclient.NewDeleteGitMergedBranchesRequest(int64(123), "DepotPath_example") // DeleteGitMergedBranchesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitMergedBranches(context.Background()).Authorization(authorization).Action(action).DeleteGitMergedBranchesRequest(deleteGitMergedBranchesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitMergedBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitMergedBranches`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitMergedBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitMergedBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitMergedBranches&quot;]
 **deleteGitMergedBranchesRequest** | [**DeleteGitMergedBranchesRequest**](DeleteGitMergedBranchesRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitProtectedTagRule

> DeleteAPIDoc200Response DeleteGitProtectedTagRule(ctx).Authorization(authorization).Action(action).DeleteGitProtectedTagRuleRequest(deleteGitProtectedTagRuleRequest).Execute()

标签信息-删除标签保护规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitProtectedTagRule" // string | Action (default to "DeleteGitProtectedTagRule")
	deleteGitProtectedTagRuleRequest := *openapiclient.NewDeleteGitProtectedTagRuleRequest(int64(123), "Rule_example") // DeleteGitProtectedTagRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitProtectedTagRule(context.Background()).Authorization(authorization).Action(action).DeleteGitProtectedTagRuleRequest(deleteGitProtectedTagRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitProtectedTagRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitProtectedTagRule`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitProtectedTagRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitProtectedTagRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitProtectedTagRule&quot;]
 **deleteGitProtectedTagRuleRequest** | [**DeleteGitProtectedTagRuleRequest**](DeleteGitProtectedTagRuleRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitRelease

> DeleteAPIDoc200Response DeleteGitRelease(ctx).Authorization(authorization).Action(action).DeleteGitReleaseRequest(deleteGitReleaseRequest).Execute()

版本信息-删除仓库版本



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitRelease" // string | Action (default to "DeleteGitRelease")
	deleteGitReleaseRequest := *openapiclient.NewDeleteGitReleaseRequest(int64(123), "TagName_example") // DeleteGitReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitRelease(context.Background()).Authorization(authorization).Action(action).DeleteGitReleaseRequest(deleteGitReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitRelease`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitRelease&quot;]
 **deleteGitReleaseRequest** | [**DeleteGitReleaseRequest**](DeleteGitReleaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitTag

> DeleteAPIDoc200Response DeleteGitTag(ctx).Authorization(authorization).Action(action).DeleteGitTagRequest(deleteGitTagRequest).Execute()

标签信息-代码仓库删除tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteGitTag" // string | Action (default to "DeleteGitTag")
	deleteGitTagRequest := *openapiclient.NewDeleteGitTagRequest(int64(123), "TagName_example") // DeleteGitTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteGitTag(context.Background()).Authorization(authorization).Action(action).DeleteGitTagRequest(deleteGitTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGitTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGitTag`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteGitTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteGitTag&quot;]
 **deleteGitTagRequest** | [**DeleteGitTagRequest**](DeleteGitTagRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssue

> DeleteAPIDoc200Response DeleteIssue(ctx).Authorization(authorization).Action(action).DeleteIssueRequest(deleteIssueRequest).Execute()

事项删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteIssue" // string | Action (default to "DeleteIssue")
	deleteIssueRequest := *openapiclient.NewDeleteIssueRequest(int64(123), "ProjectName_example") // DeleteIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteIssue(context.Background()).Authorization(authorization).Action(action).DeleteIssueRequest(deleteIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssue`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteIssue&quot;]
 **deleteIssueRequest** | [**DeleteIssueRequest**](DeleteIssueRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueBlock

> DeleteAPIDoc200Response DeleteIssueBlock(ctx).Authorization(authorization).Action(action).DeleteIssueBlockRequest(deleteIssueBlockRequest).Execute()

前置事项删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteIssueBlock" // string | Action (default to "DeleteIssueBlock")
	deleteIssueBlockRequest := *openapiclient.NewDeleteIssueBlockRequest(int64(123), int64(123), "ProjectName_example") // DeleteIssueBlockRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteIssueBlock(context.Background()).Authorization(authorization).Action(action).DeleteIssueBlockRequest(deleteIssueBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIssueBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueBlock`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIssueBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteIssueBlock&quot;]
 **deleteIssueBlockRequest** | [**DeleteIssueBlockRequest**](DeleteIssueBlockRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueModule

> DeleteAPIDoc200Response DeleteIssueModule(ctx).Authorization(authorization).Action(action).DeleteIssueModuleRequest(deleteIssueModuleRequest).Execute()

事项模块删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteIssueModule" // string | Action (default to "DeleteIssueModule")
	deleteIssueModuleRequest := *openapiclient.NewDeleteIssueModuleRequest(int64(123), "ProjectName_example") // DeleteIssueModuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteIssueModule(context.Background()).Authorization(authorization).Action(action).DeleteIssueModuleRequest(deleteIssueModuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIssueModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueModule`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIssueModule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteIssueModule&quot;]
 **deleteIssueModuleRequest** | [**DeleteIssueModuleRequest**](DeleteIssueModuleRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueWorkHours

> DeleteAPIDoc200Response DeleteIssueWorkHours(ctx).Authorization(authorization).Action(action).DeleteIssueWorkHoursRequest(deleteIssueWorkHoursRequest).Execute()

工时日志删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteIssueWorkHours" // string | Action (default to "DeleteIssueWorkHours")
	deleteIssueWorkHoursRequest := *openapiclient.NewDeleteIssueWorkHoursRequest(int64(123), "ProjectName_example", false, int32(123)) // DeleteIssueWorkHoursRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteIssueWorkHours(context.Background()).Authorization(authorization).Action(action).DeleteIssueWorkHoursRequest(deleteIssueWorkHoursRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIssueWorkHours``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueWorkHours`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIssueWorkHours`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueWorkHoursRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteIssueWorkHours&quot;]
 **deleteIssueWorkHoursRequest** | [**DeleteIssueWorkHoursRequest**](DeleteIssueWorkHoursRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIteration

> DeleteAPIDoc200Response DeleteIteration(ctx).Authorization(authorization).Action(action).DeleteIterationRequest(deleteIterationRequest).Execute()

迭代删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteIteration" // string | Action (default to "DeleteIteration")
	deleteIterationRequest := *openapiclient.NewDeleteIterationRequest(int64(123), "ProjectName_example") // DeleteIterationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteIteration(context.Background()).Authorization(authorization).Action(action).DeleteIterationRequest(deleteIterationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIteration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIteration`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIteration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteIteration&quot;]
 **deleteIterationRequest** | [**DeleteIterationRequest**](DeleteIterationRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemberSshKey

> DeleteAPIDoc200Response DeleteMemberSshKey(ctx).Authorization(authorization).Action(action).DeleteMemberSshKeyRequest(deleteMemberSshKeyRequest).Execute()

仓库设置-删除团队成员的SSH公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteMemberSshKey" // string | Action (default to "DeleteMemberSshKey")
	deleteMemberSshKeyRequest := *openapiclient.NewDeleteMemberSshKeyRequest(int64(123), int32(123)) // DeleteMemberSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteMemberSshKey(context.Background()).Authorization(authorization).Action(action).DeleteMemberSshKeyRequest(deleteMemberSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMemberSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMemberSshKey`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteMemberSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteMemberSshKey&quot;]
 **deleteMemberSshKeyRequest** | [**DeleteMemberSshKeyRequest**](DeleteMemberSshKeyRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMergeRequestNote

> DeleteMergeRequestNote200Response DeleteMergeRequestNote(ctx).Authorization(authorization).Action(action).DeleteMergeRequestNoteRequest(deleteMergeRequestNoteRequest).Execute()

合并请求-删除合并请求行评论和改动文件diff行评论



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteMergeRequestNote" // string | Action (default to "DeleteMergeRequestNote")
	deleteMergeRequestNoteRequest := *openapiclient.NewDeleteMergeRequestNoteRequest("DepotPath_example", int32(123), int32(123)) // DeleteMergeRequestNoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteMergeRequestNote(context.Background()).Authorization(authorization).Action(action).DeleteMergeRequestNoteRequest(deleteMergeRequestNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMergeRequestNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMergeRequestNote`: DeleteMergeRequestNote200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteMergeRequestNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMergeRequestNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteMergeRequestNote&quot;]
 **deleteMergeRequestNoteRequest** | [**DeleteMergeRequestNoteRequest**](DeleteMergeRequestNoteRequest.md) |  | 

### Return type

[**DeleteMergeRequestNote200Response**](DeleteMergeRequestNote200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMergeRequestReviewer

> DeleteMergeRequestReviewer200Response DeleteMergeRequestReviewer(ctx).Authorization(authorization).Action(action).DeleteMergeRequestReviewerRequest(deleteMergeRequestReviewerRequest).Execute()

合并请求-删除mr评审者



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteMergeRequestReviewer" // string | Action (default to "DeleteMergeRequestReviewer")
	deleteMergeRequestReviewerRequest := *openapiclient.NewDeleteMergeRequestReviewerRequest(int64(123), int64(123), "ReviewerGlobalKey_example") // DeleteMergeRequestReviewerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteMergeRequestReviewer(context.Background()).Authorization(authorization).Action(action).DeleteMergeRequestReviewerRequest(deleteMergeRequestReviewerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMergeRequestReviewer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMergeRequestReviewer`: DeleteMergeRequestReviewer200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteMergeRequestReviewer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMergeRequestReviewerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteMergeRequestReviewer&quot;]
 **deleteMergeRequestReviewerRequest** | [**DeleteMergeRequestReviewerRequest**](DeleteMergeRequestReviewerRequest.md) |  | 

### Return type

[**DeleteMergeRequestReviewer200Response**](DeleteMergeRequestReviewer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOneProject

> DeleteAPIDoc200Response DeleteOneProject(ctx).Authorization(authorization).Action(action).DeleteOneProjectRequest(deleteOneProjectRequest).Execute()

单个项目删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteOneProject" // string | Action (default to "DeleteOneProject")
	deleteOneProjectRequest := *openapiclient.NewDeleteOneProjectRequest(int32(123)) // DeleteOneProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteOneProject(context.Background()).Authorization(authorization).Action(action).DeleteOneProjectRequest(deleteOneProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteOneProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOneProject`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteOneProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteOneProject&quot;]
 **deleteOneProjectRequest** | [**DeleteOneProjectRequest**](DeleteOneProjectRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePoliciesById

> DeletePoliciesById200Response DeletePoliciesById(ctx).Authorization(authorization).Action(action).DeletePoliciesByIdRequest(deletePoliciesByIdRequest).Execute()

权限组批量删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeletePoliciesById" // string | Action (default to "DeletePoliciesById")
	deletePoliciesByIdRequest := *openapiclient.NewDeletePoliciesByIdRequest([]int64{int64(123)}) // DeletePoliciesByIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeletePoliciesById(context.Background()).Authorization(authorization).Action(action).DeletePoliciesByIdRequest(deletePoliciesByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePoliciesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePoliciesById`: DeletePoliciesById200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeletePoliciesById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoliciesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeletePoliciesById&quot;]
 **deletePoliciesByIdRequest** | [**DeletePoliciesByIdRequest**](DeletePoliciesByIdRequest.md) |  | 

### Return type

[**DeletePoliciesById200Response**](DeletePoliciesById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProgramMemberPolicy

> DeleteProgramMemberPolicy200Response DeleteProgramMemberPolicy(ctx).Authorization(authorization).Action(action).DeleteProgramMemberPolicyRequest(deleteProgramMemberPolicyRequest).Execute()

项目集成员权限组删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteProgramMemberPolicy" // string | Action (default to "DeleteProgramMemberPolicy")
	deleteProgramMemberPolicyRequest := *openapiclient.NewDeleteProgramMemberPolicyRequest([]string{"Policies_example"}, "11", "USER", int64(12)) // DeleteProgramMemberPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteProgramMemberPolicy(context.Background()).Authorization(authorization).Action(action).DeleteProgramMemberPolicyRequest(deleteProgramMemberPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProgramMemberPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProgramMemberPolicy`: DeleteProgramMemberPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteProgramMemberPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProgramMemberPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteProgramMemberPolicy&quot;]
 **deleteProgramMemberPolicyRequest** | [**DeleteProgramMemberPolicyRequest**](DeleteProgramMemberPolicyRequest.md) |  | 

### Return type

[**DeleteProgramMemberPolicy200Response**](DeleteProgramMemberPolicy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectAnnouncement

> DeleteAPIDoc200Response DeleteProjectAnnouncement(ctx).Authorization(authorization).Action(action).DeleteProjectAnnouncementRequest(deleteProjectAnnouncementRequest).Execute()

项目公告删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteProjectAnnouncement" // string | Action (default to "DeleteProjectAnnouncement")
	deleteProjectAnnouncementRequest := *openapiclient.NewDeleteProjectAnnouncementRequest("ProjectName_example", []int64{int64(123)}) // DeleteProjectAnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteProjectAnnouncement(context.Background()).Authorization(authorization).Action(action).DeleteProjectAnnouncementRequest(deleteProjectAnnouncementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProjectAnnouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectAnnouncement`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteProjectAnnouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteProjectAnnouncement&quot;]
 **deleteProjectAnnouncementRequest** | [**DeleteProjectAnnouncementRequest**](DeleteProjectAnnouncementRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectLabel

> DeleteProjectLabel200Response DeleteProjectLabel(ctx).Authorization(authorization).Action(action).DeleteProjectLabelRequest(deleteProjectLabelRequest).Execute()

项目标签删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteProjectLabel" // string | Action (default to "DeleteProjectLabel")
	deleteProjectLabelRequest := *openapiclient.NewDeleteProjectLabelRequest(int64(123), int64(123)) // DeleteProjectLabelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteProjectLabel(context.Background()).Authorization(authorization).Action(action).DeleteProjectLabelRequest(deleteProjectLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProjectLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectLabel`: DeleteProjectLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteProjectLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteProjectLabel&quot;]
 **deleteProjectLabelRequest** | [**DeleteProjectLabelRequest**](DeleteProjectLabelRequest.md) |  | 

### Return type

[**DeleteProjectLabel200Response**](DeleteProjectLabel200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectMemberPrincipal

> DeleteAPIDoc200Response DeleteProjectMemberPrincipal(ctx).Authorization(authorization).Action(action).DeleteProjectMemberPrincipalRequest(deleteProjectMemberPrincipalRequest).Execute()

项目成员主体删除(包含用户组、部门、成员)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteProjectMemberPrincipal" // string | Action (default to "DeleteProjectMemberPrincipal")
	deleteProjectMemberPrincipalRequest := *openapiclient.NewDeleteProjectMemberPrincipalRequest([]openapiclient.PrincipalDel{*openapiclient.NewPrincipalDel("PrincipalId_example", "PrincipalType_example")}, int32(123)) // DeleteProjectMemberPrincipalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteProjectMemberPrincipal(context.Background()).Authorization(authorization).Action(action).DeleteProjectMemberPrincipalRequest(deleteProjectMemberPrincipalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProjectMemberPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectMemberPrincipal`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteProjectMemberPrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectMemberPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteProjectMemberPrincipal&quot;]
 **deleteProjectMemberPrincipalRequest** | [**DeleteProjectMemberPrincipalRequest**](DeleteProjectMemberPrincipalRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRelease

> DeleteAPIDoc200Response DeleteRelease(ctx).Authorization(authorization).Action(action).DeleteReleaseRequest(deleteReleaseRequest).Execute()

版本删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteRelease" // string | Action (default to "DeleteRelease")
	deleteReleaseRequest := *openapiclient.NewDeleteReleaseRequest("ProjectName_example", int64(123)) // DeleteReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteRelease(context.Background()).Authorization(authorization).Action(action).DeleteReleaseRequest(deleteReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRelease`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteRelease&quot;]
 **deleteReleaseRequest** | [**DeleteReleaseRequest**](DeleteReleaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReport

> DeleteAPIDoc200Response DeleteReport(ctx).Authorization(authorization).Action(action).DeleteReportRequest(deleteReportRequest).Execute()

测试报告删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteReport" // string | Action (default to "DeleteReport")
	deleteReportRequest := *openapiclient.NewDeleteReportRequest("ProjectName_example", int32(123)) // DeleteReportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteReport(context.Background()).Authorization(authorization).Action(action).DeleteReportRequest(deleteReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReport`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteReport&quot;]
 **deleteReportRequest** | [**DeleteReportRequest**](DeleteReportRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequirementDefectRelation

> DeleteAPIDoc200Response DeleteRequirementDefectRelation(ctx).Authorization(authorization).Action(action).CreateRequirementDefectRelationRequest(createRequirementDefectRelationRequest).Execute()

需求取消关联缺陷



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteRequirementDefectRelation" // string | Action (default to "DeleteRequirementDefectRelation")
	createRequirementDefectRelationRequest := *openapiclient.NewCreateRequirementDefectRelationRequest(int64(123), "ProjectName_example", int64(123)) // CreateRequirementDefectRelationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteRequirementDefectRelation(context.Background()).Authorization(authorization).Action(action).CreateRequirementDefectRelationRequest(createRequirementDefectRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRequirementDefectRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRequirementDefectRelation`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteRequirementDefectRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequirementDefectRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteRequirementDefectRelation&quot;]
 **createRequirementDefectRelationRequest** | [**CreateRequirementDefectRelationRequest**](CreateRequirementDefectRelationRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSshKey

> DeleteAPIDoc200Response DeleteSshKey(ctx).Authorization(authorization).Action(action).DeleteSshKeyRequest(deleteSshKeyRequest).Execute()

仓库设置-删除当前用户的SSH公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteSshKey" // string | Action (default to "DeleteSshKey")
	deleteSshKeyRequest := *openapiclient.NewDeleteSshKeyRequest(int32(123)) // DeleteSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteSshKey(context.Background()).Authorization(authorization).Action(action).DeleteSshKeyRequest(deleteSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSshKey`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteSshKey&quot;]
 **deleteSshKeyRequest** | [**DeleteSshKeyRequest**](DeleteSshKeyRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamLevelDepotSpec

> DeleteAPIDoc200Response DeleteTeamLevelDepotSpec(ctx).Authorization(authorization).Action(action).DeleteTeamLevelDepotSpecRequest(deleteTeamLevelDepotSpecRequest).Execute()

仓库设置-删除团队级别的分支规范



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteTeamLevelDepotSpec" // string | Action (default to "DeleteTeamLevelDepotSpec")
	deleteTeamLevelDepotSpecRequest := *openapiclient.NewDeleteTeamLevelDepotSpecRequest("DepotSpecName_example") // DeleteTeamLevelDepotSpecRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteTeamLevelDepotSpec(context.Background()).Authorization(authorization).Action(action).DeleteTeamLevelDepotSpecRequest(deleteTeamLevelDepotSpecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTeamLevelDepotSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeamLevelDepotSpec`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTeamLevelDepotSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamLevelDepotSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteTeamLevelDepotSpec&quot;]
 **deleteTeamLevelDepotSpecRequest** | [**DeleteTeamLevelDepotSpecRequest**](DeleteTeamLevelDepotSpecRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamMember

> DeleteAPIDoc200Response DeleteTeamMember(ctx).Authorization(authorization).Action(action).DeleteTeamMemberRequest(deleteTeamMemberRequest).Execute()

团队成员删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteTeamMember" // string | Action (default to "DeleteTeamMember")
	deleteTeamMemberRequest := *openapiclient.NewDeleteTeamMemberRequest(int32(123)) // DeleteTeamMemberRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteTeamMember(context.Background()).Authorization(authorization).Action(action).DeleteTeamMemberRequest(deleteTeamMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeamMember`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTeamMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteTeamMember&quot;]
 **deleteTeamMemberRequest** | [**DeleteTeamMemberRequest**](DeleteTeamMemberRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTestCase

> DeleteAPIDoc200Response DeleteTestCase(ctx).Authorization(authorization).Action(action).DeleteTestCaseRequest(deleteTestCaseRequest).Execute()

测试用例删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteTestCase" // string | Action (default to "DeleteTestCase")
	deleteTestCaseRequest := *openapiclient.NewDeleteTestCaseRequest(int32(123), "ProjectName_example") // DeleteTestCaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteTestCase(context.Background()).Authorization(authorization).Action(action).DeleteTestCaseRequest(deleteTestCaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTestCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTestCase`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTestCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteTestCase&quot;]
 **deleteTestCaseRequest** | [**DeleteTestCaseRequest**](DeleteTestCaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTestCaseSection

> DeleteAPIDoc200Response DeleteTestCaseSection(ctx).Authorization(authorization).Action(action).DeleteTestCaseSectionRequest(deleteTestCaseSectionRequest).Execute()

测试用例分组删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteTestCaseSection" // string | Action (default to "DeleteTestCaseSection")
	deleteTestCaseSectionRequest := *openapiclient.NewDeleteTestCaseSectionRequest("ProjectName_example", int32(123)) // DeleteTestCaseSectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteTestCaseSection(context.Background()).Authorization(authorization).Action(action).DeleteTestCaseSectionRequest(deleteTestCaseSectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTestCaseSection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTestCaseSection`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTestCaseSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestCaseSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteTestCaseSection&quot;]
 **deleteTestCaseSectionRequest** | [**DeleteTestCaseSectionRequest**](DeleteTestCaseSectionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTestRun

> DeleteAPIDoc200Response DeleteTestRun(ctx).Authorization(authorization).Action(action).DeleteTestRunRequest(deleteTestRunRequest).Execute()

测试计划删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteTestRun" // string | Action (default to "DeleteTestRun")
	deleteTestRunRequest := *openapiclient.NewDeleteTestRunRequest("ProjectName_example", int32(123)) // DeleteTestRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteTestRun(context.Background()).Authorization(authorization).Action(action).DeleteTestRunRequest(deleteTestRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTestRun`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTestRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteTestRun&quot;]
 **deleteTestRunRequest** | [**DeleteTestRunRequest**](DeleteTestRunRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroupByIds

> DeleteAPIDoc200Response DeleteUserGroupByIds(ctx).Authorization(authorization).Action(action).DeleteUserGroupByIdsRequest(deleteUserGroupByIdsRequest).Execute()

用户组删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteUserGroupByIds" // string | Action (default to "DeleteUserGroupByIds")
	deleteUserGroupByIdsRequest := *openapiclient.NewDeleteUserGroupByIdsRequest([]int64{int64(123)}) // DeleteUserGroupByIdsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteUserGroupByIds(context.Background()).Authorization(authorization).Action(action).DeleteUserGroupByIdsRequest(deleteUserGroupByIdsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserGroupByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserGroupByIds`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteUserGroupByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteUserGroupByIds&quot;]
 **deleteUserGroupByIdsRequest** | [**DeleteUserGroupByIdsRequest**](DeleteUserGroupByIdsRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserGroupUsers

> DeleteAPIDoc200Response DeleteUserGroupUsers(ctx).Authorization(authorization).Action(action).DeleteUserGroupUsersRequest(deleteUserGroupUsersRequest).Execute()

用户组删除用户



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteUserGroupUsers" // string | Action (default to "DeleteUserGroupUsers")
	deleteUserGroupUsersRequest := *openapiclient.NewDeleteUserGroupUsersRequest([]openapiclient.UserGroupUserInfos{*openapiclient.NewUserGroupUserInfos(NullableInt64(123), NullableInt64(123))}) // DeleteUserGroupUsersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DeleteUserGroupUsers(context.Background()).Authorization(authorization).Action(action).DeleteUserGroupUsersRequest(deleteUserGroupUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserGroupUsers`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteUserGroupUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteUserGroupUsers&quot;]
 **deleteUserGroupUsersRequest** | [**DeleteUserGroupUsersRequest**](DeleteUserGroupUsersRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAgentSecret

> DescribeAgentSecret200Response DescribeAgentSecret(ctx).Authorization(authorization).Action(action).Body(body).Execute()

堡垒机安装 Secret



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAgentSecret" // string | Action (default to "DescribeAgentSecret")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeAgentSecret(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeAgentSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAgentSecret`: DescribeAgentSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeAgentSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAgentSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAgentSecret&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeAgentSecret200Response**](DescribeAgentSecret200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAllMergeRequestNotes

> DescribeAllMergeRequestNotes200Response DescribeAllMergeRequestNotes(ctx).Authorization(authorization).Action(action).DescribeAllMergeRequestNotesRequest(describeAllMergeRequestNotesRequest).Execute()

合并请求-获取所有合并请求行评论和改动文件diff行评论



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAllMergeRequestNotes" // string | Action (default to "DescribeAllMergeRequestNotes")
	describeAllMergeRequestNotesRequest := *openapiclient.NewDescribeAllMergeRequestNotesRequest("DepotPath_example", int32(123), int32(123)) // DescribeAllMergeRequestNotesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeAllMergeRequestNotes(context.Background()).Authorization(authorization).Action(action).DescribeAllMergeRequestNotesRequest(describeAllMergeRequestNotesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeAllMergeRequestNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAllMergeRequestNotes`: DescribeAllMergeRequestNotes200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeAllMergeRequestNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAllMergeRequestNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAllMergeRequestNotes&quot;]
 **describeAllMergeRequestNotesRequest** | [**DescribeAllMergeRequestNotesRequest**](DescribeAllMergeRequestNotesRequest.md) |  | 

### Return type

[**DescribeAllMergeRequestNotes200Response**](DescribeAllMergeRequestNotes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAllProjectLabels

> DescribeAllProjectLabels200Response DescribeAllProjectLabels(ctx).Authorization(authorization).Action(action).DescribeAllProjectLabelsRequest(describeAllProjectLabelsRequest).Execute()

项目标签查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAllProjectLabels" // string | Action (default to "DescribeAllProjectLabels")
	describeAllProjectLabelsRequest := *openapiclient.NewDescribeAllProjectLabelsRequest(int64(123)) // DescribeAllProjectLabelsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeAllProjectLabels(context.Background()).Authorization(authorization).Action(action).DescribeAllProjectLabelsRequest(describeAllProjectLabelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeAllProjectLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAllProjectLabels`: DescribeAllProjectLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeAllProjectLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAllProjectLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAllProjectLabels&quot;]
 **describeAllProjectLabelsRequest** | [**DescribeAllProjectLabelsRequest**](DescribeAllProjectLabelsRequest.md) |  | 

### Return type

[**DescribeAllProjectLabels200Response**](DescribeAllProjectLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAllProjectsIssueWorkLogList

> DescribeAllProjectsIssueWorkLogList200Response DescribeAllProjectsIssueWorkLogList(ctx).Authorization(authorization).Action(action).DescribeAllProjectsIssueWorkLogListRequest(describeAllProjectsIssueWorkLogListRequest).Execute()

工时日志列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAllProjectsIssueWorkLogList" // string | Action (default to "DescribeAllProjectsIssueWorkLogList")
	describeAllProjectsIssueWorkLogListRequest := *openapiclient.NewDescribeAllProjectsIssueWorkLogListRequest() // DescribeAllProjectsIssueWorkLogListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeAllProjectsIssueWorkLogList(context.Background()).Authorization(authorization).Action(action).DescribeAllProjectsIssueWorkLogListRequest(describeAllProjectsIssueWorkLogListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeAllProjectsIssueWorkLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAllProjectsIssueWorkLogList`: DescribeAllProjectsIssueWorkLogList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeAllProjectsIssueWorkLogList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAllProjectsIssueWorkLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAllProjectsIssueWorkLogList&quot;]
 **describeAllProjectsIssueWorkLogListRequest** | [**DescribeAllProjectsIssueWorkLogListRequest**](DescribeAllProjectsIssueWorkLogListRequest.md) |  | 

### Return type

[**DescribeAllProjectsIssueWorkLogList200Response**](DescribeAllProjectsIssueWorkLogList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactChecksums

> DescribeArtifactChecksums200Response DescribeArtifactChecksums(ctx).Authorization(authorization).Action(action).DescribeArtifactChecksumsRequest(describeArtifactChecksumsRequest).Execute()

制品Checksum列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactChecksums" // string | Action (default to "DescribeArtifactChecksums")
	describeArtifactChecksumsRequest := *openapiclient.NewDescribeArtifactChecksumsRequest("Package_example", "PackageVersion_example", int32(123), "Repository_example") // DescribeArtifactChecksumsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactChecksums(context.Background()).Authorization(authorization).Action(action).DescribeArtifactChecksumsRequest(describeArtifactChecksumsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactChecksums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactChecksums`: DescribeArtifactChecksums200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactChecksums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactChecksumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactChecksums&quot;]
 **describeArtifactChecksumsRequest** | [**DescribeArtifactChecksumsRequest**](DescribeArtifactChecksumsRequest.md) |  | 

### Return type

[**DescribeArtifactChecksums200Response**](DescribeArtifactChecksums200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactCredit

> DescribeArtifactCredit200Response DescribeArtifactCredit(ctx).Authorization(authorization).Action(action).DescribeArtifactCreditRequest(describeArtifactCreditRequest).Execute()

查询制品授信清单详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactCredit" // string | Action (default to "DescribeArtifactCredit")
	describeArtifactCreditRequest := *openapiclient.NewDescribeArtifactCreditRequest(int64(123)) // DescribeArtifactCreditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactCredit(context.Background()).Authorization(authorization).Action(action).DescribeArtifactCreditRequest(describeArtifactCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactCredit`: DescribeArtifactCredit200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactCredit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactCredit&quot;]
 **describeArtifactCreditRequest** | [**DescribeArtifactCreditRequest**](DescribeArtifactCreditRequest.md) |  | 

### Return type

[**DescribeArtifactCredit200Response**](DescribeArtifactCredit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactCreditList

> DescribeArtifactCreditList200Response DescribeArtifactCreditList(ctx).Authorization(authorization).Action(action).Body(body).Execute()

制品授信清单列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactCreditList" // string | Action (default to "DescribeArtifactCreditList")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactCreditList(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactCreditList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactCreditList`: DescribeArtifactCreditList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactCreditList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactCreditListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactCreditList&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeArtifactCreditList200Response**](DescribeArtifactCreditList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactFileDownloadUrl

> DescribeArtifactFileDownloadUrl200Response DescribeArtifactFileDownloadUrl(ctx).Authorization(authorization).Action(action).DescribeArtifactFileDownloadUrlRequest(describeArtifactFileDownloadUrlRequest).Execute()

制品文件临时下载链接获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactFileDownloadUrl" // string | Action (default to "DescribeArtifactFileDownloadUrl")
	describeArtifactFileDownloadUrlRequest := *openapiclient.NewDescribeArtifactFileDownloadUrlRequest("FileName_example", "Package_example", "PackageVersion_example", int32(123), "Repository_example") // DescribeArtifactFileDownloadUrlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactFileDownloadUrl(context.Background()).Authorization(authorization).Action(action).DescribeArtifactFileDownloadUrlRequest(describeArtifactFileDownloadUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactFileDownloadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactFileDownloadUrl`: DescribeArtifactFileDownloadUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactFileDownloadUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactFileDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactFileDownloadUrl&quot;]
 **describeArtifactFileDownloadUrlRequest** | [**DescribeArtifactFileDownloadUrlRequest**](DescribeArtifactFileDownloadUrlRequest.md) |  | 

### Return type

[**DescribeArtifactFileDownloadUrl200Response**](DescribeArtifactFileDownloadUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactPackageList

> DescribeArtifactPackageList200Response DescribeArtifactPackageList(ctx).Authorization(authorization).Action(action).DescribeArtifactPackageListRequest(describeArtifactPackageListRequest).Execute()

制品包（镜像）列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactPackageList" // string | Action (default to "DescribeArtifactPackageList")
	describeArtifactPackageListRequest := *openapiclient.NewDescribeArtifactPackageListRequest(int32(123), "Repository_example") // DescribeArtifactPackageListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactPackageList(context.Background()).Authorization(authorization).Action(action).DescribeArtifactPackageListRequest(describeArtifactPackageListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactPackageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactPackageList`: DescribeArtifactPackageList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactPackageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactPackageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactPackageList&quot;]
 **describeArtifactPackageListRequest** | [**DescribeArtifactPackageListRequest**](DescribeArtifactPackageListRequest.md) |  | 

### Return type

[**DescribeArtifactPackageList200Response**](DescribeArtifactPackageList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactProperties

> DescribeArtifactProperties200Response DescribeArtifactProperties(ctx).Authorization(authorization).Action(action).DescribeArtifactChecksumsRequest(describeArtifactChecksumsRequest).Execute()

制品属性列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactProperties" // string | Action (default to "DescribeArtifactProperties")
	describeArtifactChecksumsRequest := *openapiclient.NewDescribeArtifactChecksumsRequest("Package_example", "PackageVersion_example", int32(123), "Repository_example") // DescribeArtifactChecksumsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactProperties(context.Background()).Authorization(authorization).Action(action).DescribeArtifactChecksumsRequest(describeArtifactChecksumsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactProperties`: DescribeArtifactProperties200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactProperties&quot;]
 **describeArtifactChecksumsRequest** | [**DescribeArtifactChecksumsRequest**](DescribeArtifactChecksumsRequest.md) |  | 

### Return type

[**DescribeArtifactProperties200Response**](DescribeArtifactProperties200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactRepositoryFileList

> DescribeArtifactRepositoryFileList200Response DescribeArtifactRepositoryFileList(ctx).Authorization(authorization).Action(action).DescribeArtifactRepositoryFileListRequest(describeArtifactRepositoryFileListRequest).Execute()

制品仓库下可下载的文件列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactRepositoryFileList" // string | Action (default to "DescribeArtifactRepositoryFileList")
	describeArtifactRepositoryFileListRequest := *openapiclient.NewDescribeArtifactRepositoryFileListRequest("Project_example", "Repository_example") // DescribeArtifactRepositoryFileListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactRepositoryFileList(context.Background()).Authorization(authorization).Action(action).DescribeArtifactRepositoryFileListRequest(describeArtifactRepositoryFileListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactRepositoryFileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactRepositoryFileList`: DescribeArtifactRepositoryFileList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactRepositoryFileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactRepositoryFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactRepositoryFileList&quot;]
 **describeArtifactRepositoryFileListRequest** | [**DescribeArtifactRepositoryFileListRequest**](DescribeArtifactRepositoryFileListRequest.md) |  | 

### Return type

[**DescribeArtifactRepositoryFileList200Response**](DescribeArtifactRepositoryFileList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactRepositoryList

> DescribeArtifactRepositoryList200Response DescribeArtifactRepositoryList(ctx).Authorization(authorization).Action(action).DescribeArtifactRepositoryListRequest(describeArtifactRepositoryListRequest).Execute()

制品仓库列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactRepositoryList" // string | Action (default to "DescribeArtifactRepositoryList")
	describeArtifactRepositoryListRequest := *openapiclient.NewDescribeArtifactRepositoryListRequest(int32(123)) // DescribeArtifactRepositoryListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactRepositoryList(context.Background()).Authorization(authorization).Action(action).DescribeArtifactRepositoryListRequest(describeArtifactRepositoryListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactRepositoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactRepositoryList`: DescribeArtifactRepositoryList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactRepositoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactRepositoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactRepositoryList&quot;]
 **describeArtifactRepositoryListRequest** | [**DescribeArtifactRepositoryListRequest**](DescribeArtifactRepositoryListRequest.md) |  | 

### Return type

[**DescribeArtifactRepositoryList200Response**](DescribeArtifactRepositoryList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactVersionFileList

> DescribeArtifactVersionFileList200Response DescribeArtifactVersionFileList(ctx).Authorization(authorization).Action(action).DescribeArtifactVersionFileListRequest(describeArtifactVersionFileListRequest).Execute()

制品版本可下载的文件列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactVersionFileList" // string | Action (default to "DescribeArtifactVersionFileList")
	describeArtifactVersionFileListRequest := *openapiclient.NewDescribeArtifactVersionFileListRequest("Package_example", "PackageVersion_example", int32(123), "Repository_example") // DescribeArtifactVersionFileListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactVersionFileList(context.Background()).Authorization(authorization).Action(action).DescribeArtifactVersionFileListRequest(describeArtifactVersionFileListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactVersionFileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactVersionFileList`: DescribeArtifactVersionFileList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactVersionFileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactVersionFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactVersionFileList&quot;]
 **describeArtifactVersionFileListRequest** | [**DescribeArtifactVersionFileListRequest**](DescribeArtifactVersionFileListRequest.md) |  | 

### Return type

[**DescribeArtifactVersionFileList200Response**](DescribeArtifactVersionFileList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeArtifactVersionList

> DescribeArtifactVersionList200Response DescribeArtifactVersionList(ctx).Authorization(authorization).Action(action).DescribeArtifactVersionListRequest(describeArtifactVersionListRequest).Execute()

制品版本列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeArtifactVersionList" // string | Action (default to "DescribeArtifactVersionList")
	describeArtifactVersionListRequest := *openapiclient.NewDescribeArtifactVersionListRequest("Package_example", int32(123), "Repository_example") // DescribeArtifactVersionListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeArtifactVersionList(context.Background()).Authorization(authorization).Action(action).DescribeArtifactVersionListRequest(describeArtifactVersionListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeArtifactVersionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeArtifactVersionList`: DescribeArtifactVersionList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeArtifactVersionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeArtifactVersionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeArtifactVersionList&quot;]
 **describeArtifactVersionListRequest** | [**DescribeArtifactVersionListRequest**](DescribeArtifactVersionListRequest.md) |  | 

### Return type

[**DescribeArtifactVersionList200Response**](DescribeArtifactVersionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAvailablePoliciesOnResource

> DescribeAvailablePoliciesOnResource200Response DescribeAvailablePoliciesOnResource(ctx).Authorization(authorization).Action(action).DescribeAvailablePoliciesOnResourceRequest(describeAvailablePoliciesOnResourceRequest).Execute()

权限组列表查询（指定资源）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAvailablePoliciesOnResource" // string | Action (default to "DescribeAvailablePoliciesOnResource")
	describeAvailablePoliciesOnResourceRequest := *openapiclient.NewDescribeAvailablePoliciesOnResourceRequest(*openapiclient.NewDescribeAvailablePoliciesOnResourceRequestFilter(), int32(123), int32(123), *openapiclient.NewResourceInfoOfPolicyScope("ResourceId_example", "ResourceType_example")) // DescribeAvailablePoliciesOnResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeAvailablePoliciesOnResource(context.Background()).Authorization(authorization).Action(action).DescribeAvailablePoliciesOnResourceRequest(describeAvailablePoliciesOnResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeAvailablePoliciesOnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAvailablePoliciesOnResource`: DescribeAvailablePoliciesOnResource200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeAvailablePoliciesOnResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAvailablePoliciesOnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAvailablePoliciesOnResource&quot;]
 **describeAvailablePoliciesOnResourceRequest** | [**DescribeAvailablePoliciesOnResourceRequest**](DescribeAvailablePoliciesOnResourceRequest.md) |  | 

### Return type

[**DescribeAvailablePoliciesOnResource200Response**](DescribeAvailablePoliciesOnResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeBlockIssueList

> DescribeBlockIssueList200Response DescribeBlockIssueList(ctx).Authorization(authorization).Action(action).DescribeBlockIssueListRequest(describeBlockIssueListRequest).Execute()

后置事项查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeBlockIssueList" // string | Action (default to "DescribeBlockIssueList")
	describeBlockIssueListRequest := *openapiclient.NewDescribeBlockIssueListRequest(int64(123), "ProjectName_example") // DescribeBlockIssueListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeBlockIssueList(context.Background()).Authorization(authorization).Action(action).DescribeBlockIssueListRequest(describeBlockIssueListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeBlockIssueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeBlockIssueList`: DescribeBlockIssueList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeBlockIssueList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeBlockIssueListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeBlockIssueList&quot;]
 **describeBlockIssueListRequest** | [**DescribeBlockIssueListRequest**](DescribeBlockIssueListRequest.md) |  | 

### Return type

[**DescribeBlockIssueList200Response**](DescribeBlockIssueList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeBlockedByIssueList

> DescribeBlockedByIssueList200Response DescribeBlockedByIssueList(ctx).Authorization(authorization).Action(action).DescribeIssueCommentListRequest(describeIssueCommentListRequest).Execute()

前置事项查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeBlockedByIssueList" // string | Action (default to "DescribeBlockedByIssueList")
	describeIssueCommentListRequest := *openapiclient.NewDescribeIssueCommentListRequest(int64(123), "ProjectName_example") // DescribeIssueCommentListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeBlockedByIssueList(context.Background()).Authorization(authorization).Action(action).DescribeIssueCommentListRequest(describeIssueCommentListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeBlockedByIssueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeBlockedByIssueList`: DescribeBlockedByIssueList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeBlockedByIssueList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeBlockedByIssueListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeBlockedByIssueList&quot;]
 **describeIssueCommentListRequest** | [**DescribeIssueCommentListRequest**](DescribeIssueCommentListRequest.md) |  | 

### Return type

[**DescribeBlockedByIssueList200Response**](DescribeBlockedByIssueList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeBranchProtection

> DescribeBranchProtection200Response DescribeBranchProtection(ctx).Authorization(authorization).Action(action).DescribeBranchProtectionRequest(describeBranchProtectionRequest).Execute()

仓库设置-查询单个保护分支规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeBranchProtection" // string | Action (default to "DescribeBranchProtection")
	describeBranchProtectionRequest := *openapiclient.NewDescribeBranchProtectionRequest(int64(123), int64(123)) // DescribeBranchProtectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeBranchProtection(context.Background()).Authorization(authorization).Action(action).DescribeBranchProtectionRequest(describeBranchProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeBranchProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeBranchProtection`: DescribeBranchProtection200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeBranchProtection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeBranchProtection&quot;]
 **describeBranchProtectionRequest** | [**DescribeBranchProtectionRequest**](DescribeBranchProtectionRequest.md) |  | 

### Return type

[**DescribeBranchProtection200Response**](DescribeBranchProtection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeBranchProtectionMembers

> DescribeBranchProtectionMembers200Response DescribeBranchProtectionMembers(ctx).Authorization(authorization).Action(action).DescribeBranchProtectionMembersRequest(describeBranchProtectionMembersRequest).Execute()

仓库设置-查询保护分支规则下所有管理员信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeBranchProtectionMembers" // string | Action (default to "DescribeBranchProtectionMembers")
	describeBranchProtectionMembersRequest := *openapiclient.NewDescribeBranchProtectionMembersRequest(int64(123), int64(123)) // DescribeBranchProtectionMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeBranchProtectionMembers(context.Background()).Authorization(authorization).Action(action).DescribeBranchProtectionMembersRequest(describeBranchProtectionMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeBranchProtectionMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeBranchProtectionMembers`: DescribeBranchProtectionMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeBranchProtectionMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeBranchProtectionMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeBranchProtectionMembers&quot;]
 **describeBranchProtectionMembersRequest** | [**DescribeBranchProtectionMembersRequest**](DescribeBranchProtectionMembersRequest.md) |  | 

### Return type

[**DescribeBranchProtectionMembers200Response**](DescribeBranchProtectionMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeBranchProtections

> DescribeBranchProtections200Response DescribeBranchProtections(ctx).Authorization(authorization).Action(action).DescribeBranchProtectionsRequest(describeBranchProtectionsRequest).Execute()

仓库设置-查询仓库保护分支规则集合



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeBranchProtections" // string | Action (default to "DescribeBranchProtections")
	describeBranchProtectionsRequest := *openapiclient.NewDescribeBranchProtectionsRequest(int64(123)) // DescribeBranchProtectionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeBranchProtections(context.Background()).Authorization(authorization).Action(action).DescribeBranchProtectionsRequest(describeBranchProtectionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeBranchProtections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeBranchProtections`: DescribeBranchProtections200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeBranchProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeBranchProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeBranchProtections&quot;]
 **describeBranchProtectionsRequest** | [**DescribeBranchProtectionsRequest**](DescribeBranchProtectionsRequest.md) |  | 

### Return type

[**DescribeBranchProtections200Response**](DescribeBranchProtections200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCanMerge

> DescribeCanMerge200Response DescribeCanMerge(ctx).Authorization(authorization).Action(action).DescribeCanMergeRequest(describeCanMergeRequest).Execute()

合并请求-查看两个分支是否可以合并



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCanMerge" // string | Action (default to "DescribeCanMerge")
	describeCanMergeRequest := *openapiclient.NewDescribeCanMergeRequest(int64(123), "Source_example", "Target_example") // DescribeCanMergeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCanMerge(context.Background()).Authorization(authorization).Action(action).DescribeCanMergeRequest(describeCanMergeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCanMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCanMerge`: DescribeCanMerge200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCanMerge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCanMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCanMerge&quot;]
 **describeCanMergeRequest** | [**DescribeCanMergeRequest**](DescribeCanMergeRequest.md) |  | 

### Return type

[**DescribeCanMerge200Response**](DescribeCanMerge200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdAgentMachines

> DescribeCdAgentMachines200Response DescribeCdAgentMachines(ctx).Authorization(authorization).Action(action).DescribeCdAgentMachinesRequest(describeCdAgentMachinesRequest).Execute()

CD 堡垒机列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdAgentMachines" // string | Action (default to "DescribeCdAgentMachines")
	describeCdAgentMachinesRequest := *openapiclient.NewDescribeCdAgentMachinesRequest() // DescribeCdAgentMachinesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdAgentMachines(context.Background()).Authorization(authorization).Action(action).DescribeCdAgentMachinesRequest(describeCdAgentMachinesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdAgentMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdAgentMachines`: DescribeCdAgentMachines200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdAgentMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdAgentMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdAgentMachines&quot;]
 **describeCdAgentMachinesRequest** | [**DescribeCdAgentMachinesRequest**](DescribeCdAgentMachinesRequest.md) |  | 

### Return type

[**DescribeCdAgentMachines200Response**](DescribeCdAgentMachines200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdApplication

> DescribeCdApplication200Response DescribeCdApplication(ctx).Authorization(authorization).Action(action).DescribeCdApplicationRequest(describeCdApplicationRequest).Execute()

CD 应用详情获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdApplication" // string | Action (default to "DescribeCdApplication")
	describeCdApplicationRequest := *openapiclient.NewDescribeCdApplicationRequest("Application_example") // DescribeCdApplicationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdApplication(context.Background()).Authorization(authorization).Action(action).DescribeCdApplicationRequest(describeCdApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdApplication`: DescribeCdApplication200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdApplication&quot;]
 **describeCdApplicationRequest** | [**DescribeCdApplicationRequest**](DescribeCdApplicationRequest.md) |  | 

### Return type

[**DescribeCdApplication200Response**](DescribeCdApplication200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdApplications

> DescribeCdApplications200Response DescribeCdApplications(ctx).Authorization(authorization).Action(action).Body(body).Execute()

CD 应用列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdApplications" // string | Action (default to "DescribeCdApplications")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdApplications(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdApplications`: DescribeCdApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdApplications&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeCdApplications200Response**](DescribeCdApplications200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdApplicationsByProject

> DescribeCdApplicationsByProject200Response DescribeCdApplicationsByProject(ctx).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()

关联应用列表获取（指定项目名）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdApplicationsByProject" // string | Action (default to "DescribeCdApplicationsByProject")
	describeAPIDocListRequest := *openapiclient.NewDescribeAPIDocListRequest("ProjectName_example") // DescribeAPIDocListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdApplicationsByProject(context.Background()).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdApplicationsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdApplicationsByProject`: DescribeCdApplicationsByProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdApplicationsByProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdApplicationsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdApplicationsByProject&quot;]
 **describeAPIDocListRequest** | [**DescribeAPIDocListRequest**](DescribeAPIDocListRequest.md) |  | 

### Return type

[**DescribeCdApplicationsByProject200Response**](DescribeCdApplicationsByProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdCloudAccounts

> DescribeCdCloudAccounts200Response DescribeCdCloudAccounts(ctx).Authorization(authorization).Action(action).DescribeCdCloudAccountsRequest(describeCdCloudAccountsRequest).Execute()

CD 云账号列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdCloudAccounts" // string | Action (default to "DescribeCdCloudAccounts")
	describeCdCloudAccountsRequest := *openapiclient.NewDescribeCdCloudAccountsRequest() // DescribeCdCloudAccountsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdCloudAccounts(context.Background()).Authorization(authorization).Action(action).DescribeCdCloudAccountsRequest(describeCdCloudAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdCloudAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdCloudAccounts`: DescribeCdCloudAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdCloudAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdCloudAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdCloudAccounts&quot;]
 **describeCdCloudAccountsRequest** | [**DescribeCdCloudAccountsRequest**](DescribeCdCloudAccountsRequest.md) |  | 

### Return type

[**DescribeCdCloudAccounts200Response**](DescribeCdCloudAccounts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdDeployCountByApplications

> DescribeCdDeployCountByApplications200Response DescribeCdDeployCountByApplications(ctx).Authorization(authorization).Action(action).DescribeCdDeployCountByApplicationsRequest(describeCdDeployCountByApplicationsRequest).Execute()

发布次数-根据应用名列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdDeployCountByApplications" // string | Action (default to "DescribeCdDeployCountByApplications")
	describeCdDeployCountByApplicationsRequest := *openapiclient.NewDescribeCdDeployCountByApplicationsRequest([]string{"Application_example"}, "EndAt_example", "StartAt_example") // DescribeCdDeployCountByApplicationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdDeployCountByApplications(context.Background()).Authorization(authorization).Action(action).DescribeCdDeployCountByApplicationsRequest(describeCdDeployCountByApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdDeployCountByApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdDeployCountByApplications`: DescribeCdDeployCountByApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdDeployCountByApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdDeployCountByApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdDeployCountByApplications&quot;]
 **describeCdDeployCountByApplicationsRequest** | [**DescribeCdDeployCountByApplicationsRequest**](DescribeCdDeployCountByApplicationsRequest.md) |  | 

### Return type

[**DescribeCdDeployCountByApplications200Response**](DescribeCdDeployCountByApplications200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdDeployCountByProject

> DescribeCdDeployCountByProject200Response DescribeCdDeployCountByProject(ctx).Authorization(authorization).Action(action).DescribeCdDeployTimeByProjectRequest(describeCdDeployTimeByProjectRequest).Execute()

关联应用的发布次数获取（指定项目名）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdDeployCountByProject" // string | Action (default to "DescribeCdDeployCountByProject")
	describeCdDeployTimeByProjectRequest := *openapiclient.NewDescribeCdDeployTimeByProjectRequest("EndAt_example", "ProjectName_example", "StartAt_example") // DescribeCdDeployTimeByProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdDeployCountByProject(context.Background()).Authorization(authorization).Action(action).DescribeCdDeployTimeByProjectRequest(describeCdDeployTimeByProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdDeployCountByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdDeployCountByProject`: DescribeCdDeployCountByProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdDeployCountByProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdDeployCountByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdDeployCountByProject&quot;]
 **describeCdDeployTimeByProjectRequest** | [**DescribeCdDeployTimeByProjectRequest**](DescribeCdDeployTimeByProjectRequest.md) |  | 

### Return type

[**DescribeCdDeployCountByProject200Response**](DescribeCdDeployCountByProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdDeployTimeByApplications

> DescribeCdDeployTimeByApplications200Response DescribeCdDeployTimeByApplications(ctx).Authorization(authorization).Action(action).DescribeCdDeployTimeByApplicationsRequest(describeCdDeployTimeByApplicationsRequest).Execute()

发布时长-根据应用名列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdDeployTimeByApplications" // string | Action (default to "DescribeCdDeployTimeByApplications")
	describeCdDeployTimeByApplicationsRequest := *openapiclient.NewDescribeCdDeployTimeByApplicationsRequest([]string{"Application_example"}, "EndAt_example", "StartAt_example") // DescribeCdDeployTimeByApplicationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdDeployTimeByApplications(context.Background()).Authorization(authorization).Action(action).DescribeCdDeployTimeByApplicationsRequest(describeCdDeployTimeByApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdDeployTimeByApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdDeployTimeByApplications`: DescribeCdDeployTimeByApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdDeployTimeByApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdDeployTimeByApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdDeployTimeByApplications&quot;]
 **describeCdDeployTimeByApplicationsRequest** | [**DescribeCdDeployTimeByApplicationsRequest**](DescribeCdDeployTimeByApplicationsRequest.md) |  | 

### Return type

[**DescribeCdDeployTimeByApplications200Response**](DescribeCdDeployTimeByApplications200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdDeployTimeByProject

> DescribeCdDeployTimeByProject200Response DescribeCdDeployTimeByProject(ctx).Authorization(authorization).Action(action).DescribeCdDeployTimeByProjectRequest(describeCdDeployTimeByProjectRequest).Execute()

关联应用的发布时长-根据项目名获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdDeployTimeByProject" // string | Action (default to "DescribeCdDeployTimeByProject")
	describeCdDeployTimeByProjectRequest := *openapiclient.NewDescribeCdDeployTimeByProjectRequest("EndAt_example", "ProjectName_example", "StartAt_example") // DescribeCdDeployTimeByProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdDeployTimeByProject(context.Background()).Authorization(authorization).Action(action).DescribeCdDeployTimeByProjectRequest(describeCdDeployTimeByProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdDeployTimeByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdDeployTimeByProject`: DescribeCdDeployTimeByProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdDeployTimeByProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdDeployTimeByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdDeployTimeByProject&quot;]
 **describeCdDeployTimeByProjectRequest** | [**DescribeCdDeployTimeByProjectRequest**](DescribeCdDeployTimeByProjectRequest.md) |  | 

### Return type

[**DescribeCdDeployTimeByProject200Response**](DescribeCdDeployTimeByProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdDeployTrendByApplications

> DescribeCdDeployTrendByApplications200Response DescribeCdDeployTrendByApplications(ctx).Authorization(authorization).Action(action).DescribeCdDeployTrendByApplicationsRequest(describeCdDeployTrendByApplicationsRequest).Execute()

发布趋势-根据应用名列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdDeployTrendByApplications" // string | Action (default to "DescribeCdDeployTrendByApplications")
	describeCdDeployTrendByApplicationsRequest := *openapiclient.NewDescribeCdDeployTrendByApplicationsRequest([]string{"Application_example"}, "EndAt_example", "StartAt_example") // DescribeCdDeployTrendByApplicationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdDeployTrendByApplications(context.Background()).Authorization(authorization).Action(action).DescribeCdDeployTrendByApplicationsRequest(describeCdDeployTrendByApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdDeployTrendByApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdDeployTrendByApplications`: DescribeCdDeployTrendByApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdDeployTrendByApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdDeployTrendByApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdDeployTrendByApplications&quot;]
 **describeCdDeployTrendByApplicationsRequest** | [**DescribeCdDeployTrendByApplicationsRequest**](DescribeCdDeployTrendByApplicationsRequest.md) |  | 

### Return type

[**DescribeCdDeployTrendByApplications200Response**](DescribeCdDeployTrendByApplications200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdDeployTrendByProject

> DescribeCdDeployTrendByProject200Response DescribeCdDeployTrendByProject(ctx).Authorization(authorization).Action(action).DescribeCdDeployTrendByProjectRequest(describeCdDeployTrendByProjectRequest).Execute()

关联应用的发布趋势-根据项目名获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdDeployTrendByProject" // string | Action (default to "DescribeCdDeployTrendByProject")
	describeCdDeployTrendByProjectRequest := *openapiclient.NewDescribeCdDeployTrendByProjectRequest("EndAt_example", "ProjectName_example", "StartAt_example") // DescribeCdDeployTrendByProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdDeployTrendByProject(context.Background()).Authorization(authorization).Action(action).DescribeCdDeployTrendByProjectRequest(describeCdDeployTrendByProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdDeployTrendByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdDeployTrendByProject`: DescribeCdDeployTrendByProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdDeployTrendByProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdDeployTrendByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdDeployTrendByProject&quot;]
 **describeCdDeployTrendByProjectRequest** | [**DescribeCdDeployTrendByProjectRequest**](DescribeCdDeployTrendByProjectRequest.md) |  | 

### Return type

[**DescribeCdDeployTrendByProject200Response**](DescribeCdDeployTrendByProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdHostServerGroup

> DescribeCdHostServerGroup200Response DescribeCdHostServerGroup(ctx).Authorization(authorization).Action(action).DescribeCdHostServerGroupRequest(describeCdHostServerGroupRequest).Execute()

CD 主机组获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdHostServerGroup" // string | Action (default to "DescribeCdHostServerGroup")
	describeCdHostServerGroupRequest := *openapiclient.NewDescribeCdHostServerGroupRequest(int64(123)) // DescribeCdHostServerGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdHostServerGroup(context.Background()).Authorization(authorization).Action(action).DescribeCdHostServerGroupRequest(describeCdHostServerGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdHostServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdHostServerGroup`: DescribeCdHostServerGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdHostServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdHostServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdHostServerGroup&quot;]
 **describeCdHostServerGroupRequest** | [**DescribeCdHostServerGroupRequest**](DescribeCdHostServerGroupRequest.md) |  | 

### Return type

[**DescribeCdHostServerGroup200Response**](DescribeCdHostServerGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdHostServerGroups

> DescribeCdHostServerGroups200Response DescribeCdHostServerGroups(ctx).Authorization(authorization).Action(action).DescribeCdHostServerGroupsRequest(describeCdHostServerGroupsRequest).Execute()

CD 主机组列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdHostServerGroups" // string | Action (default to "DescribeCdHostServerGroups")
	describeCdHostServerGroupsRequest := *openapiclient.NewDescribeCdHostServerGroupsRequest() // DescribeCdHostServerGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdHostServerGroups(context.Background()).Authorization(authorization).Action(action).DescribeCdHostServerGroupsRequest(describeCdHostServerGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdHostServerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdHostServerGroups`: DescribeCdHostServerGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdHostServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdHostServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdHostServerGroups&quot;]
 **describeCdHostServerGroupsRequest** | [**DescribeCdHostServerGroupsRequest**](DescribeCdHostServerGroupsRequest.md) |  | 

### Return type

[**DescribeCdHostServerGroups200Response**](DescribeCdHostServerGroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdPipeline

> DescribeCdPipeline200Response DescribeCdPipeline(ctx).Authorization(authorization).Action(action).DescribeCdPipelineRequest(describeCdPipelineRequest).Execute()

CD 部署流程执行记录获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdPipeline" // string | Action (default to "DescribeCdPipeline")
	describeCdPipelineRequest := *openapiclient.NewDescribeCdPipelineRequest("PipelineExecutionId_example") // DescribeCdPipelineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdPipeline(context.Background()).Authorization(authorization).Action(action).DescribeCdPipelineRequest(describeCdPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdPipeline`: DescribeCdPipeline200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdPipeline&quot;]
 **describeCdPipelineRequest** | [**DescribeCdPipelineRequest**](DescribeCdPipelineRequest.md) |  | 

### Return type

[**DescribeCdPipeline200Response**](DescribeCdPipeline200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdPipelineConfig

> DescribeCdPipelineConfig200Response DescribeCdPipelineConfig(ctx).Authorization(authorization).Action(action).DescribeCdPipelineConfigRequest(describeCdPipelineConfigRequest).Execute()

CD 部署流程配置-根据名称获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdPipelineConfig" // string | Action (default to "DescribeCdPipelineConfig")
	describeCdPipelineConfigRequest := *openapiclient.NewDescribeCdPipelineConfigRequest("Application_example", "PipelineName_example") // DescribeCdPipelineConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdPipelineConfig(context.Background()).Authorization(authorization).Action(action).DescribeCdPipelineConfigRequest(describeCdPipelineConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdPipelineConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdPipelineConfig`: DescribeCdPipelineConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdPipelineConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdPipelineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdPipelineConfig&quot;]
 **describeCdPipelineConfigRequest** | [**DescribeCdPipelineConfigRequest**](DescribeCdPipelineConfigRequest.md) |  | 

### Return type

[**DescribeCdPipelineConfig200Response**](DescribeCdPipelineConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdPipelineConfigs

> DescribeCdPipelineConfigs200Response DescribeCdPipelineConfigs(ctx).Authorization(authorization).Action(action).DescribeCdPipelineConfigsRequest(describeCdPipelineConfigsRequest).Execute()

CD 应用下的所有部署流程配置获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdPipelineConfigs" // string | Action (default to "DescribeCdPipelineConfigs")
	describeCdPipelineConfigsRequest := *openapiclient.NewDescribeCdPipelineConfigsRequest("Application_example") // DescribeCdPipelineConfigsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdPipelineConfigs(context.Background()).Authorization(authorization).Action(action).DescribeCdPipelineConfigsRequest(describeCdPipelineConfigsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdPipelineConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdPipelineConfigs`: DescribeCdPipelineConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdPipelineConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdPipelineConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdPipelineConfigs&quot;]
 **describeCdPipelineConfigsRequest** | [**DescribeCdPipelineConfigsRequest**](DescribeCdPipelineConfigsRequest.md) |  | 

### Return type

[**DescribeCdPipelineConfigs200Response**](DescribeCdPipelineConfigs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCdTask

> DescribeCdTask200Response DescribeCdTask(ctx).Authorization(authorization).Action(action).DescribeCdTaskRequest(describeCdTaskRequest).Execute()

CD 任务执行记录获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCdTask" // string | Action (default to "DescribeCdTask")
	describeCdTaskRequest := *openapiclient.NewDescribeCdTaskRequest("TaskExecutionId_example") // DescribeCdTaskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCdTask(context.Background()).Authorization(authorization).Action(action).DescribeCdTaskRequest(describeCdTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCdTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCdTask`: DescribeCdTask200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCdTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCdTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCdTask&quot;]
 **describeCdTaskRequest** | [**DescribeCdTaskRequest**](DescribeCdTaskRequest.md) |  | 

### Return type

[**DescribeCdTask200Response**](DescribeCdTask200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodeSearch

> DescribeCodeSearch200Response DescribeCodeSearch(ctx).Authorization(authorization).Action(action).DescribeCodeSearchRequest(describeCodeSearchRequest).Execute()

仓库信息-查询代码片段详细列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodeSearch" // string | Action (default to "DescribeCodeSearch")
	describeCodeSearchRequest := *openapiclient.NewDescribeCodeSearchRequest("BranchName_example", "DepotPath_example", "KeyWord_example", int64(123), int64(123)) // DescribeCodeSearchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodeSearch(context.Background()).Authorization(authorization).Action(action).DescribeCodeSearchRequest(describeCodeSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodeSearch`: DescribeCodeSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodeSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodeSearch&quot;]
 **describeCodeSearchRequest** | [**DescribeCodeSearchRequest**](DescribeCodeSearchRequest.md) |  | 

### Return type

[**DescribeCodeSearch200Response**](DescribeCodeSearch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuild

> DescribeCodingCIBuild200Response DescribeCodingCIBuild(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()

构建记录详情查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuild" // string | Action (default to "DescribeCodingCIBuild")
	describeCodingCIBuildStageRequest := *openapiclient.NewDescribeCodingCIBuildStageRequest(int32(123)) // DescribeCodingCIBuildStageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuild(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuild`: DescribeCodingCIBuild200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuild&quot;]
 **describeCodingCIBuildStageRequest** | [**DescribeCodingCIBuildStageRequest**](DescribeCodingCIBuildStageRequest.md) |  | 

### Return type

[**DescribeCodingCIBuild200Response**](DescribeCodingCIBuild200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildArtifacts

> DescribeCodingCIBuildArtifacts200Response DescribeCodingCIBuildArtifacts(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildArtifactsRequest(describeCodingCIBuildArtifactsRequest).Execute()

构建任务制品查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildArtifacts" // string | Action (default to "DescribeCodingCIBuildArtifacts")
	describeCodingCIBuildArtifactsRequest := *openapiclient.NewDescribeCodingCIBuildArtifactsRequest(int64(123)) // DescribeCodingCIBuildArtifactsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildArtifacts(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildArtifactsRequest(describeCodingCIBuildArtifactsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildArtifacts`: DescribeCodingCIBuildArtifacts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildArtifacts&quot;]
 **describeCodingCIBuildArtifactsRequest** | [**DescribeCodingCIBuildArtifactsRequest**](DescribeCodingCIBuildArtifactsRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildArtifacts200Response**](DescribeCodingCIBuildArtifacts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildEnvs

> DescribeCodingCIBuildEnvs200Response DescribeCodingCIBuildEnvs(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildEnvsRequest(describeCodingCIBuildEnvsRequest).Execute()

构建计划环境变量获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildEnvs" // string | Action (default to "DescribeCodingCIBuildEnvs")
	describeCodingCIBuildEnvsRequest := *openapiclient.NewDescribeCodingCIBuildEnvsRequest(int32(123), "Type_example") // DescribeCodingCIBuildEnvsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildEnvs(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildEnvsRequest(describeCodingCIBuildEnvsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildEnvs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildEnvs`: DescribeCodingCIBuildEnvs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildEnvs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildEnvsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildEnvs&quot;]
 **describeCodingCIBuildEnvsRequest** | [**DescribeCodingCIBuildEnvsRequest**](DescribeCodingCIBuildEnvsRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildEnvs200Response**](DescribeCodingCIBuildEnvs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildHtmlReports

> DescribeCodingCIBuildHtmlReports200Response DescribeCodingCIBuildHtmlReports(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildArtifactsRequest(describeCodingCIBuildArtifactsRequest).Execute()

构建任务网页报告查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildHtmlReports" // string | Action (default to "DescribeCodingCIBuildHtmlReports")
	describeCodingCIBuildArtifactsRequest := *openapiclient.NewDescribeCodingCIBuildArtifactsRequest(int64(123)) // DescribeCodingCIBuildArtifactsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildHtmlReports(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildArtifactsRequest(describeCodingCIBuildArtifactsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildHtmlReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildHtmlReports`: DescribeCodingCIBuildHtmlReports200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildHtmlReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildHtmlReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildHtmlReports&quot;]
 **describeCodingCIBuildArtifactsRequest** | [**DescribeCodingCIBuildArtifactsRequest**](DescribeCodingCIBuildArtifactsRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildHtmlReports200Response**](DescribeCodingCIBuildHtmlReports200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildLog

> DescribeCodingCIBuildLog200Response DescribeCodingCIBuildLog(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildLogRequest(describeCodingCIBuildLogRequest).Execute()

构建日志获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildLog" // string | Action (default to "DescribeCodingCIBuildLog")
	describeCodingCIBuildLogRequest := *openapiclient.NewDescribeCodingCIBuildLogRequest(int32(123), int32(123)) // DescribeCodingCIBuildLogRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildLog(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildLogRequest(describeCodingCIBuildLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildLog`: DescribeCodingCIBuildLog200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildLog&quot;]
 **describeCodingCIBuildLogRequest** | [**DescribeCodingCIBuildLogRequest**](DescribeCodingCIBuildLogRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildLog200Response**](DescribeCodingCIBuildLog200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildLogRaw

> DescribeCodingCIBuildLogRaw200Response DescribeCodingCIBuildLogRaw(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildLogRawRequest(describeCodingCIBuildLogRawRequest).Execute()

构建完整日志查询（原始日志 Raw）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildLogRaw" // string | Action (default to "DescribeCodingCIBuildLogRaw")
	describeCodingCIBuildLogRawRequest := *openapiclient.NewDescribeCodingCIBuildLogRawRequest(int32(123), int32(123)) // DescribeCodingCIBuildLogRawRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildLogRaw(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildLogRawRequest(describeCodingCIBuildLogRawRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildLogRaw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildLogRaw`: DescribeCodingCIBuildLogRaw200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildLogRaw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildLogRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildLogRaw&quot;]
 **describeCodingCIBuildLogRawRequest** | [**DescribeCodingCIBuildLogRawRequest**](DescribeCodingCIBuildLogRawRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildLogRaw200Response**](DescribeCodingCIBuildLogRaw200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildMetrics

> DescribeCodingCIBuildMetrics200Response DescribeCodingCIBuildMetrics(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildMetricsRequest(describeCodingCIBuildMetricsRequest).Execute()

构建计划度量查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildMetrics" // string | Action (default to "DescribeCodingCIBuildMetrics")
	describeCodingCIBuildMetricsRequest := *openapiclient.NewDescribeCodingCIBuildMetricsRequest("EndTime_example", "StartTime_example") // DescribeCodingCIBuildMetricsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildMetrics(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildMetricsRequest(describeCodingCIBuildMetricsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildMetrics`: DescribeCodingCIBuildMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildMetrics&quot;]
 **describeCodingCIBuildMetricsRequest** | [**DescribeCodingCIBuildMetricsRequest**](DescribeCodingCIBuildMetricsRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildMetrics200Response**](DescribeCodingCIBuildMetrics200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildStage

> DescribeCodingCIBuildStage200Response DescribeCodingCIBuildStage(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()

构建任务阶段获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildStage" // string | Action (default to "DescribeCodingCIBuildStage")
	describeCodingCIBuildStageRequest := *openapiclient.NewDescribeCodingCIBuildStageRequest(int32(123)) // DescribeCodingCIBuildStageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildStage(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildStage`: DescribeCodingCIBuildStage200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildStage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildStage&quot;]
 **describeCodingCIBuildStageRequest** | [**DescribeCodingCIBuildStageRequest**](DescribeCodingCIBuildStageRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildStage200Response**](DescribeCodingCIBuildStage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildStatistics

> DescribeCodingCIBuildStatistics200Response DescribeCodingCIBuildStatistics(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStatisticsRequest(describeCodingCIBuildStatisticsRequest).Execute()

构建任务统计



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildStatistics" // string | Action (default to "DescribeCodingCIBuildStatistics")
	describeCodingCIBuildStatisticsRequest := *openapiclient.NewDescribeCodingCIBuildStatisticsRequest("EndTime_example", "MetricType_example", int32(123), "StartTime_example") // DescribeCodingCIBuildStatisticsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildStatistics(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStatisticsRequest(describeCodingCIBuildStatisticsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildStatistics`: DescribeCodingCIBuildStatistics200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildStatistics&quot;]
 **describeCodingCIBuildStatisticsRequest** | [**DescribeCodingCIBuildStatisticsRequest**](DescribeCodingCIBuildStatisticsRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildStatistics200Response**](DescribeCodingCIBuildStatistics200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildStep

> DescribeCodingCIBuildStep200Response DescribeCodingCIBuildStep(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStepRequest(describeCodingCIBuildStepRequest).Execute()

构建任务指定阶段的步骤获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildStep" // string | Action (default to "DescribeCodingCIBuildStep")
	describeCodingCIBuildStepRequest := *openapiclient.NewDescribeCodingCIBuildStepRequest(int32(123), int32(123)) // DescribeCodingCIBuildStepRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildStep(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStepRequest(describeCodingCIBuildStepRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildStep`: DescribeCodingCIBuildStep200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildStep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildStep&quot;]
 **describeCodingCIBuildStepRequest** | [**DescribeCodingCIBuildStepRequest**](DescribeCodingCIBuildStepRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildStep200Response**](DescribeCodingCIBuildStep200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuildStepLog

> DescribeCodingCIBuildStepLog200Response DescribeCodingCIBuildStepLog(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStepLogRequest(describeCodingCIBuildStepLogRequest).Execute()

构建步骤日志获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuildStepLog" // string | Action (default to "DescribeCodingCIBuildStepLog")
	describeCodingCIBuildStepLogRequest := *openapiclient.NewDescribeCodingCIBuildStepLogRequest(int32(123), int32(123), int32(123), int32(123)) // DescribeCodingCIBuildStepLogRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuildStepLog(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStepLogRequest(describeCodingCIBuildStepLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuildStepLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuildStepLog`: DescribeCodingCIBuildStepLog200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuildStepLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildStepLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuildStepLog&quot;]
 **describeCodingCIBuildStepLogRequest** | [**DescribeCodingCIBuildStepLogRequest**](DescribeCodingCIBuildStepLogRequest.md) |  | 

### Return type

[**DescribeCodingCIBuildStepLog200Response**](DescribeCodingCIBuildStepLog200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIBuilds

> DescribeCodingCIBuilds200Response DescribeCodingCIBuilds(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildsRequest(describeCodingCIBuildsRequest).Execute()

构建计划的构建列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIBuilds" // string | Action (default to "DescribeCodingCIBuilds")
	describeCodingCIBuildsRequest := *openapiclient.NewDescribeCodingCIBuildsRequest(int32(123), int32(123), int32(123)) // DescribeCodingCIBuildsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIBuilds(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildsRequest(describeCodingCIBuildsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIBuilds`: DescribeCodingCIBuilds200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIBuilds&quot;]
 **describeCodingCIBuildsRequest** | [**DescribeCodingCIBuildsRequest**](DescribeCodingCIBuildsRequest.md) |  | 

### Return type

[**DescribeCodingCIBuilds200Response**](DescribeCodingCIBuilds200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIJob

> DescribeCodingCIJob200Response DescribeCodingCIJob(ctx).Authorization(authorization).Action(action).DeleteCodingCIJobRequest(deleteCodingCIJobRequest).Execute()

构建计划详情获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIJob" // string | Action (default to "DescribeCodingCIJob")
	deleteCodingCIJobRequest := *openapiclient.NewDeleteCodingCIJobRequest(int32(123)) // DeleteCodingCIJobRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIJob(context.Background()).Authorization(authorization).Action(action).DeleteCodingCIJobRequest(deleteCodingCIJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIJob`: DescribeCodingCIJob200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIJob&quot;]
 **deleteCodingCIJobRequest** | [**DeleteCodingCIJobRequest**](DeleteCodingCIJobRequest.md) |  | 

### Return type

[**DescribeCodingCIJob200Response**](DescribeCodingCIJob200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCIJobs

> DescribeCodingCIJobs200Response DescribeCodingCIJobs(ctx).Authorization(authorization).Action(action).DescribeCodingCIJobsRequest(describeCodingCIJobsRequest).Execute()

构建计划查询（通过项目ID）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCIJobs" // string | Action (default to "DescribeCodingCIJobs")
	describeCodingCIJobsRequest := *openapiclient.NewDescribeCodingCIJobsRequest(int32(123)) // DescribeCodingCIJobsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCIJobs(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIJobsRequest(describeCodingCIJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCIJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCIJobs`: DescribeCodingCIJobs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCIJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCIJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCIJobs&quot;]
 **describeCodingCIJobsRequest** | [**DescribeCodingCIJobsRequest**](DescribeCodingCIJobsRequest.md) |  | 

### Return type

[**DescribeCodingCIJobs200Response**](DescribeCodingCIJobs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingCurrentUser

> DescribeCodingCurrentUser200Response DescribeCodingCurrentUser(ctx).Authorization(authorization).Action(action).Body(body).Execute()

当前用户信息查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingCurrentUser" // string | Action (default to "DescribeCodingCurrentUser")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingCurrentUser(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingCurrentUser`: DescribeCodingCurrentUser200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingCurrentUser&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeCodingCurrentUser200Response**](DescribeCodingCurrentUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCodingProjects

> DescribeCodingProjects200Response DescribeCodingProjects(ctx).Authorization(authorization).Action(action).DescribeCodingProjectsRequest(describeCodingProjectsRequest).Execute()

项目列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCodingProjects" // string | Action (default to "DescribeCodingProjects")
	describeCodingProjectsRequest := *openapiclient.NewDescribeCodingProjectsRequest(int32(123), int32(123)) // DescribeCodingProjectsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCodingProjects(context.Background()).Authorization(authorization).Action(action).DescribeCodingProjectsRequest(describeCodingProjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCodingProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCodingProjects`: DescribeCodingProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCodingProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCodingProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCodingProjects&quot;]
 **describeCodingProjectsRequest** | [**DescribeCodingProjectsRequest**](DescribeCodingProjectsRequest.md) |  | 

### Return type

[**DescribeCodingProjects200Response**](DescribeCodingProjects200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCommitRefs

> DescribeCommitRefs200Response DescribeCommitRefs(ctx).Authorization(authorization).Action(action).DescribeCommitRefsRequest(describeCommitRefsRequest).Execute()

Git提交-查询commit的ref信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCommitRefs" // string | Action (default to "DescribeCommitRefs")
	describeCommitRefsRequest := *openapiclient.NewDescribeCommitRefsRequest("DepotPath_example", "Sha_example", "Type_example") // DescribeCommitRefsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCommitRefs(context.Background()).Authorization(authorization).Action(action).DescribeCommitRefsRequest(describeCommitRefsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCommitRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCommitRefs`: DescribeCommitRefs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCommitRefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCommitRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCommitRefs&quot;]
 **describeCommitRefsRequest** | [**DescribeCommitRefsRequest**](DescribeCommitRefsRequest.md) |  | 

### Return type

[**DescribeCommitRefs200Response**](DescribeCommitRefs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCommitsBetweenCommitAndCommit

> DescribeCommitsBetweenCommitAndCommit200Response DescribeCommitsBetweenCommitAndCommit(ctx).Authorization(authorization).Action(action).DescribeCommitsBetweenCommitAndCommitRequest(describeCommitsBetweenCommitAndCommitRequest).Execute()

Git提交-查询两个请求之间的请求列表（source target顺序正常）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeCommitsBetweenCommitAndCommit" // string | Action (default to "DescribeCommitsBetweenCommitAndCommit")
	describeCommitsBetweenCommitAndCommitRequest := *openapiclient.NewDescribeCommitsBetweenCommitAndCommitRequest(int64(123), "Source_example", "Target_example") // DescribeCommitsBetweenCommitAndCommitRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeCommitsBetweenCommitAndCommit(context.Background()).Authorization(authorization).Action(action).DescribeCommitsBetweenCommitAndCommitRequest(describeCommitsBetweenCommitAndCommitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeCommitsBetweenCommitAndCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCommitsBetweenCommitAndCommit`: DescribeCommitsBetweenCommitAndCommit200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeCommitsBetweenCommitAndCommit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCommitsBetweenCommitAndCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeCommitsBetweenCommitAndCommit&quot;]
 **describeCommitsBetweenCommitAndCommitRequest** | [**DescribeCommitsBetweenCommitAndCommitRequest**](DescribeCommitsBetweenCommitAndCommitRequest.md) |  | 

### Return type

[**DescribeCommitsBetweenCommitAndCommit200Response**](DescribeCommitsBetweenCommitAndCommit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeConfigTemplateList

> DescribeConfigTemplateList200Response DescribeConfigTemplateList(ctx).Authorization(authorization).Action(action).DescribeConfigTemplateListRequest(describeConfigTemplateListRequest).Execute()

配置方案获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeConfigTemplateList" // string | Action (default to "DescribeConfigTemplateList")
	describeConfigTemplateListRequest := *openapiclient.NewDescribeConfigTemplateListRequest(int64(123), int64(123), "ProjectName_example") // DescribeConfigTemplateListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeConfigTemplateList(context.Background()).Authorization(authorization).Action(action).DescribeConfigTemplateListRequest(describeConfigTemplateListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeConfigTemplateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeConfigTemplateList`: DescribeConfigTemplateList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeConfigTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeConfigTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeConfigTemplateList&quot;]
 **describeConfigTemplateListRequest** | [**DescribeConfigTemplateListRequest**](DescribeConfigTemplateListRequest.md) |  | 

### Return type

[**DescribeConfigTemplateList200Response**](DescribeConfigTemplateList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepartment

> DescribeDepartment200Response DescribeDepartment(ctx).Authorization(authorization).Action(action).DescribeDepartmentRequest(describeDepartmentRequest).Execute()

部门详情查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepartment" // string | Action (default to "DescribeDepartment")
	describeDepartmentRequest := *openapiclient.NewDescribeDepartmentRequest(int64(123)) // DescribeDepartmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepartment(context.Background()).Authorization(authorization).Action(action).DescribeDepartmentRequest(describeDepartmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepartment`: DescribeDepartment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepartment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepartment&quot;]
 **describeDepartmentRequest** | [**DescribeDepartmentRequest**](DescribeDepartmentRequest.md) |  | 

### Return type

[**DescribeDepartment200Response**](DescribeDepartment200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepartmentMembers

> DescribeDepartmentMembers200Response DescribeDepartmentMembers(ctx).Authorization(authorization).Action(action).DescribeDepartmentMembersRequest(describeDepartmentMembersRequest).Execute()

部门成员列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepartmentMembers" // string | Action (default to "DescribeDepartmentMembers")
	describeDepartmentMembersRequest := *openapiclient.NewDescribeDepartmentMembersRequest(int64(123), int64(123), int64(123)) // DescribeDepartmentMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepartmentMembers(context.Background()).Authorization(authorization).Action(action).DescribeDepartmentMembersRequest(describeDepartmentMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepartmentMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepartmentMembers`: DescribeDepartmentMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepartmentMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepartmentMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepartmentMembers&quot;]
 **describeDepartmentMembersRequest** | [**DescribeDepartmentMembersRequest**](DescribeDepartmentMembersRequest.md) |  | 

### Return type

[**DescribeDepartmentMembers200Response**](DescribeDepartmentMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotByNameInfo

> DescribeDepotByNameInfo200Response DescribeDepotByNameInfo(ctx).Authorization(authorization).Action(action).DescribeDepotByNameInfoRequest(describeDepotByNameInfoRequest).Execute()

仓库信息-查询项目下所有的仓库信息列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotByNameInfo" // string | Action (default to "DescribeDepotByNameInfo")
	describeDepotByNameInfoRequest := *openapiclient.NewDescribeDepotByNameInfoRequest("DepotName_example", "ProjectName_example", "TeamGk_example") // DescribeDepotByNameInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotByNameInfo(context.Background()).Authorization(authorization).Action(action).DescribeDepotByNameInfoRequest(describeDepotByNameInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotByNameInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotByNameInfo`: DescribeDepotByNameInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotByNameInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotByNameInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotByNameInfo&quot;]
 **describeDepotByNameInfoRequest** | [**DescribeDepotByNameInfoRequest**](DescribeDepotByNameInfoRequest.md) |  | 

### Return type

[**DescribeDepotByNameInfo200Response**](DescribeDepotByNameInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotDefaultBranch

> DescribeDepotDefaultBranch200Response DescribeDepotDefaultBranch(ctx).Authorization(authorization).Action(action).DescribeDepotDefaultBranchRequest(describeDepotDefaultBranchRequest).Execute()

仓库分支-查询仓库的默认分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotDefaultBranch" // string | Action (default to "DescribeDepotDefaultBranch")
	describeDepotDefaultBranchRequest := *openapiclient.NewDescribeDepotDefaultBranchRequest("DepotPath_example") // DescribeDepotDefaultBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotDefaultBranch(context.Background()).Authorization(authorization).Action(action).DescribeDepotDefaultBranchRequest(describeDepotDefaultBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotDefaultBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotDefaultBranch`: DescribeDepotDefaultBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotDefaultBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotDefaultBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotDefaultBranch&quot;]
 **describeDepotDefaultBranchRequest** | [**DescribeDepotDefaultBranchRequest**](DescribeDepotDefaultBranchRequest.md) |  | 

### Return type

[**DescribeDepotDefaultBranch200Response**](DescribeDepotDefaultBranch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotFilePushRules

> DescribeDepotFilePushRules200Response DescribeDepotFilePushRules(ctx).Authorization(authorization).Action(action).DescribeDepotFilePushRulesRequest(describeDepotFilePushRulesRequest).Execute()

仓库设置-查询git仓库文件推送规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotFilePushRules" // string | Action (default to "DescribeDepotFilePushRules")
	describeDepotFilePushRulesRequest := *openapiclient.NewDescribeDepotFilePushRulesRequest("DepotPath_example") // DescribeDepotFilePushRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotFilePushRules(context.Background()).Authorization(authorization).Action(action).DescribeDepotFilePushRulesRequest(describeDepotFilePushRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotFilePushRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotFilePushRules`: DescribeDepotFilePushRules200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotFilePushRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotFilePushRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotFilePushRules&quot;]
 **describeDepotFilePushRulesRequest** | [**DescribeDepotFilePushRulesRequest**](DescribeDepotFilePushRulesRequest.md) |  | 

### Return type

[**DescribeDepotFilePushRules200Response**](DescribeDepotFilePushRules200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotMergeRequests

> DescribeDepotMergeRequests200Response DescribeDepotMergeRequests(ctx).Authorization(authorization).Action(action).DescribeDepotMergeRequestsRequest(describeDepotMergeRequestsRequest).Execute()

合并请求-查询仓库合并请求列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotMergeRequests" // string | Action (default to "DescribeDepotMergeRequests")
	describeDepotMergeRequestsRequest := *openapiclient.NewDescribeDepotMergeRequestsRequest(int64(123)) // DescribeDepotMergeRequestsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotMergeRequests(context.Background()).Authorization(authorization).Action(action).DescribeDepotMergeRequestsRequest(describeDepotMergeRequestsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotMergeRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotMergeRequests`: DescribeDepotMergeRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotMergeRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotMergeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotMergeRequests&quot;]
 **describeDepotMergeRequestsRequest** | [**DescribeDepotMergeRequestsRequest**](DescribeDepotMergeRequestsRequest.md) |  | 

### Return type

[**DescribeDepotMergeRequests200Response**](DescribeDepotMergeRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotPushSetting

> DescribeDepotPushSetting200Response DescribeDepotPushSetting(ctx).Authorization(authorization).Action(action).DescribeDepotFilePushRulesRequest(describeDepotFilePushRulesRequest).Execute()

仓库设置-查询仓库推送设置



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotPushSetting" // string | Action (default to "DescribeDepotPushSetting")
	describeDepotFilePushRulesRequest := *openapiclient.NewDescribeDepotFilePushRulesRequest("DepotPath_example") // DescribeDepotFilePushRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotPushSetting(context.Background()).Authorization(authorization).Action(action).DescribeDepotFilePushRulesRequest(describeDepotFilePushRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotPushSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotPushSetting`: DescribeDepotPushSetting200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotPushSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotPushSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotPushSetting&quot;]
 **describeDepotFilePushRulesRequest** | [**DescribeDepotFilePushRulesRequest**](DescribeDepotFilePushRulesRequest.md) |  | 

### Return type

[**DescribeDepotPushSetting200Response**](DescribeDepotPushSetting200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotSpecDetail

> ModifyChooseDepotSpec200Response DescribeDepotSpecDetail(ctx).Authorization(authorization).Action(action).DescribeDepotSpecDetailRequest(describeDepotSpecDetailRequest).Execute()

仓库设置-查询仓库规范详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotSpecDetail" // string | Action (default to "DescribeDepotSpecDetail")
	describeDepotSpecDetailRequest := *openapiclient.NewDescribeDepotSpecDetailRequest() // DescribeDepotSpecDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotSpecDetail(context.Background()).Authorization(authorization).Action(action).DescribeDepotSpecDetailRequest(describeDepotSpecDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotSpecDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotSpecDetail`: ModifyChooseDepotSpec200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotSpecDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotSpecDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotSpecDetail&quot;]
 **describeDepotSpecDetailRequest** | [**DescribeDepotSpecDetailRequest**](DescribeDepotSpecDetailRequest.md) |  | 

### Return type

[**ModifyChooseDepotSpec200Response**](ModifyChooseDepotSpec200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDepotSpecs

> DescribeDepotSpecs200Response DescribeDepotSpecs(ctx).Authorization(authorization).Action(action).Body(body).Execute()

仓库设置-查询仓库规范列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDepotSpecs" // string | Action (default to "DescribeDepotSpecs")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDepotSpecs(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDepotSpecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDepotSpecs`: DescribeDepotSpecs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDepotSpecs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDepotSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDepotSpecs&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeDepotSpecs200Response**](DescribeDepotSpecs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDifferentBetween2Commits

> DescribeDifferentBetween2Commits200Response DescribeDifferentBetween2Commits(ctx).Authorization(authorization).Action(action).DescribeDifferentBetween2CommitsRequest(describeDifferentBetween2CommitsRequest).Execute()

Git提交-两次提交之间的文件差异（source target顺序正常）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDifferentBetween2Commits" // string | Action (default to "DescribeDifferentBetween2Commits")
	describeDifferentBetween2CommitsRequest := *openapiclient.NewDescribeDifferentBetween2CommitsRequest(int64(123), "Path_example", "Source_example", "Target_example") // DescribeDifferentBetween2CommitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDifferentBetween2Commits(context.Background()).Authorization(authorization).Action(action).DescribeDifferentBetween2CommitsRequest(describeDifferentBetween2CommitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDifferentBetween2Commits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDifferentBetween2Commits`: DescribeDifferentBetween2Commits200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDifferentBetween2Commits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDifferentBetween2CommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDifferentBetween2Commits&quot;]
 **describeDifferentBetween2CommitsRequest** | [**DescribeDifferentBetween2CommitsRequest**](DescribeDifferentBetween2CommitsRequest.md) |  | 

### Return type

[**DescribeDifferentBetween2Commits200Response**](DescribeDifferentBetween2Commits200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDifferentBetweenTwoCommits

> DescribeDifferentBetweenTwoCommits200Response DescribeDifferentBetweenTwoCommits(ctx).Authorization(authorization).Action(action).DescribeDifferentBetweenTwoCommitsRequest(describeDifferentBetweenTwoCommitsRequest).Execute()

Git提交-获取两次commit之间的文件差异详情(废弃，source target顺序不一致)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeDifferentBetweenTwoCommits" // string | Action (default to "DescribeDifferentBetweenTwoCommits")
	describeDifferentBetweenTwoCommitsRequest := *openapiclient.NewDescribeDifferentBetweenTwoCommitsRequest(int64(123), "Path_example", "Source_example", "Target_example") // DescribeDifferentBetweenTwoCommitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeDifferentBetweenTwoCommits(context.Background()).Authorization(authorization).Action(action).DescribeDifferentBetweenTwoCommitsRequest(describeDifferentBetweenTwoCommitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeDifferentBetweenTwoCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDifferentBetweenTwoCommits`: DescribeDifferentBetweenTwoCommits200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeDifferentBetweenTwoCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDifferentBetweenTwoCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeDifferentBetweenTwoCommits&quot;]
 **describeDifferentBetweenTwoCommitsRequest** | [**DescribeDifferentBetweenTwoCommitsRequest**](DescribeDifferentBetweenTwoCommitsRequest.md) |  | 

### Return type

[**DescribeDifferentBetweenTwoCommits200Response**](DescribeDifferentBetweenTwoCommits200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBlameInfo

> DescribeGitBlameInfo200Response DescribeGitBlameInfo(ctx).Authorization(authorization).Action(action).DescribeGitBlameInfoRequest(describeGitBlameInfoRequest).Execute()

Git提交-获取指定commit下某文件指定代码行的最后一次提交



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBlameInfo" // string | Action (default to "DescribeGitBlameInfo")
	describeGitBlameInfoRequest := *openapiclient.NewDescribeGitBlameInfoRequest("CommitSha_example", "DepotPath_example", "FilePath_example") // DescribeGitBlameInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBlameInfo(context.Background()).Authorization(authorization).Action(action).DescribeGitBlameInfoRequest(describeGitBlameInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBlameInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBlameInfo`: DescribeGitBlameInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBlameInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBlameInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBlameInfo&quot;]
 **describeGitBlameInfoRequest** | [**DescribeGitBlameInfoRequest**](DescribeGitBlameInfoRequest.md) |  | 

### Return type

[**DescribeGitBlameInfo200Response**](DescribeGitBlameInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBlob

> DescribeGitBlob200Response DescribeGitBlob(ctx).Authorization(authorization).Action(action).DescribeGitBlobRequest(describeGitBlobRequest).Execute()

Git文件-查询GitBlob



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBlob" // string | Action (default to "DescribeGitBlob")
	describeGitBlobRequest := *openapiclient.NewDescribeGitBlobRequest("BlobSha_example", int64(123)) // DescribeGitBlobRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBlob(context.Background()).Authorization(authorization).Action(action).DescribeGitBlobRequest(describeGitBlobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBlob`: DescribeGitBlob200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBlob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBlob&quot;]
 **describeGitBlobRequest** | [**DescribeGitBlobRequest**](DescribeGitBlobRequest.md) |  | 

### Return type

[**DescribeGitBlob200Response**](DescribeGitBlob200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBlobRaw

> DescribeGitBlobRaw200Response DescribeGitBlobRaw(ctx).Authorization(authorization).Action(action).DescribeGitBlobRawRequest(describeGitBlobRawRequest).Execute()

Git文件-查询Git Blob raw信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBlobRaw" // string | Action (default to "DescribeGitBlobRaw")
	describeGitBlobRawRequest := *openapiclient.NewDescribeGitBlobRawRequest("BlobSha_example", int64(123)) // DescribeGitBlobRawRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBlobRaw(context.Background()).Authorization(authorization).Action(action).DescribeGitBlobRawRequest(describeGitBlobRawRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBlobRaw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBlobRaw`: DescribeGitBlobRaw200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBlobRaw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBlobRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBlobRaw&quot;]
 **describeGitBlobRawRequest** | [**DescribeGitBlobRawRequest**](DescribeGitBlobRawRequest.md) |  | 

### Return type

[**DescribeGitBlobRaw200Response**](DescribeGitBlobRaw200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBranch

> DescribeGitBranch200Response DescribeGitBranch(ctx).Authorization(authorization).Action(action).DescribeGitBranchRequest(describeGitBranchRequest).Execute()

仓库分支-查询代码仓库单个分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBranch" // string | Action (default to "DescribeGitBranch")
	describeGitBranchRequest := *openapiclient.NewDescribeGitBranchRequest("BranchName_example", int64(123)) // DescribeGitBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBranch(context.Background()).Authorization(authorization).Action(action).DescribeGitBranchRequest(describeGitBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBranch`: DescribeGitBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBranch&quot;]
 **describeGitBranchRequest** | [**DescribeGitBranchRequest**](DescribeGitBranchRequest.md) |  | 

### Return type

[**DescribeGitBranch200Response**](DescribeGitBranch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBranchList

> DescribeGitBranchList200Response DescribeGitBranchList(ctx).Authorization(authorization).Action(action).DescribeGitBranchListRequest(describeGitBranchListRequest).Execute()

仓库分支-查询仓库分支列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBranchList" // string | Action (default to "DescribeGitBranchList")
	describeGitBranchListRequest := *openapiclient.NewDescribeGitBranchListRequest("DepotPath_example", int64(123), int64(123)) // DescribeGitBranchListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBranchList(context.Background()).Authorization(authorization).Action(action).DescribeGitBranchListRequest(describeGitBranchListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBranchList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBranchList`: DescribeGitBranchList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBranchList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBranchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBranchList&quot;]
 **describeGitBranchListRequest** | [**DescribeGitBranchListRequest**](DescribeGitBranchListRequest.md) |  | 

### Return type

[**DescribeGitBranchList200Response**](DescribeGitBranchList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBranches

> DescribeGitBranches200Response DescribeGitBranches(ctx).Authorization(authorization).Action(action).DescribeGitBranchesRequest(describeGitBranchesRequest).Execute()

仓库分支-查询仓库下所有分支列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBranches" // string | Action (default to "DescribeGitBranches")
	describeGitBranchesRequest := *openapiclient.NewDescribeGitBranchesRequest(int64(123)) // DescribeGitBranchesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBranches(context.Background()).Authorization(authorization).Action(action).DescribeGitBranchesRequest(describeGitBranchesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBranches`: DescribeGitBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBranches&quot;]
 **describeGitBranchesRequest** | [**DescribeGitBranchesRequest**](DescribeGitBranchesRequest.md) |  | 

### Return type

[**DescribeGitBranches200Response**](DescribeGitBranches200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitBranchesBySha

> DescribeGitBranchesBySha200Response DescribeGitBranchesBySha(ctx).Authorization(authorization).Action(action).DescribeGitBranchesByShaRequest(describeGitBranchesByShaRequest).Execute()

仓库分支-根据sha值查询所在分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitBranchesBySha" // string | Action (default to "DescribeGitBranchesBySha")
	describeGitBranchesByShaRequest := *openapiclient.NewDescribeGitBranchesByShaRequest(int64(123), "Sha_example") // DescribeGitBranchesByShaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitBranchesBySha(context.Background()).Authorization(authorization).Action(action).DescribeGitBranchesByShaRequest(describeGitBranchesByShaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitBranchesBySha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitBranchesBySha`: DescribeGitBranchesBySha200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitBranchesBySha`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitBranchesByShaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitBranchesBySha&quot;]
 **describeGitBranchesByShaRequest** | [**DescribeGitBranchesByShaRequest**](DescribeGitBranchesByShaRequest.md) |  | 

### Return type

[**DescribeGitBranchesBySha200Response**](DescribeGitBranchesBySha200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitComments

> DescribeGitCommitComments200Response DescribeGitCommitComments(ctx).Authorization(authorization).Action(action).DescribeGitCommitCommentsRequest(describeGitCommitCommentsRequest).Execute()

Git提交-获取commit评论



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitComments" // string | Action (default to "DescribeGitCommitComments")
	describeGitCommitCommentsRequest := *openapiclient.NewDescribeGitCommitCommentsRequest("DepotPath_example", int64(123), int64(123), "Sha_example") // DescribeGitCommitCommentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitComments(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitCommentsRequest(describeGitCommitCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitComments`: DescribeGitCommitComments200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitComments&quot;]
 **describeGitCommitCommentsRequest** | [**DescribeGitCommitCommentsRequest**](DescribeGitCommitCommentsRequest.md) |  | 

### Return type

[**DescribeGitCommitComments200Response**](DescribeGitCommitComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitDiff

> DescribeGitCommitDiff200Response DescribeGitCommitDiff(ctx).Authorization(authorization).Action(action).DescribeGitCommitDiffRequest(describeGitCommitDiffRequest).Execute()

Git提交-查询某次提交的diff信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitDiff" // string | Action (default to "DescribeGitCommitDiff")
	describeGitCommitDiffRequest := *openapiclient.NewDescribeGitCommitDiffRequest(int64(123), "Sha_example") // DescribeGitCommitDiffRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitDiff(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitDiffRequest(describeGitCommitDiffRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitDiff`: DescribeGitCommitDiff200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitDiff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitDiff&quot;]
 **describeGitCommitDiffRequest** | [**DescribeGitCommitDiffRequest**](DescribeGitCommitDiffRequest.md) |  | 

### Return type

[**DescribeGitCommitDiff200Response**](DescribeGitCommitDiff200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitFilePathList

> DescribeGitCommitFilePathList200Response DescribeGitCommitFilePathList(ctx).Authorization(authorization).Action(action).DescribeGitCommitFilePathListRequest(describeGitCommitFilePathListRequest).Execute()

Git提交-查询仓库某次提交改动的文件路径列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitFilePathList" // string | Action (default to "DescribeGitCommitFilePathList")
	describeGitCommitFilePathListRequest := *openapiclient.NewDescribeGitCommitFilePathListRequest("CommitSha_example", "DepotPath_example") // DescribeGitCommitFilePathListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitFilePathList(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitFilePathListRequest(describeGitCommitFilePathListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitFilePathList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitFilePathList`: DescribeGitCommitFilePathList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitFilePathList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitFilePathListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitFilePathList&quot;]
 **describeGitCommitFilePathListRequest** | [**DescribeGitCommitFilePathListRequest**](DescribeGitCommitFilePathListRequest.md) |  | 

### Return type

[**DescribeGitCommitFilePathList200Response**](DescribeGitCommitFilePathList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitInfo

> DescribeGitCommitInfo200Response DescribeGitCommitInfo(ctx).Authorization(authorization).Action(action).DescribeGitCommitInfoRequest(describeGitCommitInfoRequest).Execute()

Git提交-查询单个请求详情信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitInfo" // string | Action (default to "DescribeGitCommitInfo")
	describeGitCommitInfoRequest := *openapiclient.NewDescribeGitCommitInfoRequest(int64(123), "Sha_example") // DescribeGitCommitInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitInfo(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitInfoRequest(describeGitCommitInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitInfo`: DescribeGitCommitInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitInfo&quot;]
 **describeGitCommitInfoRequest** | [**DescribeGitCommitInfoRequest**](DescribeGitCommitInfoRequest.md) |  | 

### Return type

[**DescribeGitCommitInfo200Response**](DescribeGitCommitInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitInfos

> DescribeGitCommitInfos200Response DescribeGitCommitInfos(ctx).Authorization(authorization).Action(action).DescribeGitCommitsInPageRequest(describeGitCommitsInPageRequest).Execute()

Git提交-查询仓库分支下提交列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitInfos" // string | Action (default to "DescribeGitCommitInfos")
	describeGitCommitsInPageRequest := *openapiclient.NewDescribeGitCommitsInPageRequest(int64(123), "Ref_example") // DescribeGitCommitsInPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitInfos(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitsInPageRequest(describeGitCommitsInPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitInfos`: DescribeGitCommitInfos200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitInfos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitInfos&quot;]
 **describeGitCommitsInPageRequest** | [**DescribeGitCommitsInPageRequest**](DescribeGitCommitsInPageRequest.md) |  | 

### Return type

[**DescribeGitCommitInfos200Response**](DescribeGitCommitInfos200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitNote

> DescribeGitCommitNote200Response DescribeGitCommitNote(ctx).Authorization(authorization).Action(action).DescribeGitCommitNoteRequest(describeGitCommitNoteRequest).Execute()

Git提交-获取提交注释



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitNote" // string | Action (default to "DescribeGitCommitNote")
	describeGitCommitNoteRequest := *openapiclient.NewDescribeGitCommitNoteRequest("CommitSha_example", int64(123), "NotesRef_example") // DescribeGitCommitNoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitNote(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitNoteRequest(describeGitCommitNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitNote`: DescribeGitCommitNote200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitNote&quot;]
 **describeGitCommitNoteRequest** | [**DescribeGitCommitNoteRequest**](DescribeGitCommitNoteRequest.md) |  | 

### Return type

[**DescribeGitCommitNote200Response**](DescribeGitCommitNote200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitStatus

> DescribeGitCommitStatus200Response DescribeGitCommitStatus(ctx).Authorization(authorization).Action(action).DescribeGitCommitStatusRequest(describeGitCommitStatusRequest).Execute()

Git提交-查询提交对应的流水线状态



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitStatus" // string | Action (default to "DescribeGitCommitStatus")
	describeGitCommitStatusRequest := *openapiclient.NewDescribeGitCommitStatusRequest("CommitSha_example", int64(123)) // DescribeGitCommitStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitStatus(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitStatusRequest(describeGitCommitStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitStatus`: DescribeGitCommitStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitStatus&quot;]
 **describeGitCommitStatusRequest** | [**DescribeGitCommitStatusRequest**](DescribeGitCommitStatusRequest.md) |  | 

### Return type

[**DescribeGitCommitStatus200Response**](DescribeGitCommitStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitCommitsInPage

> DescribeGitCommitsInPage200Response DescribeGitCommitsInPage(ctx).Authorization(authorization).Action(action).DescribeGitCommitsInPageRequest(describeGitCommitsInPageRequest).Execute()

Git提交-查询仓库分支下提交列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitCommitsInPage" // string | Action (default to "DescribeGitCommitsInPage")
	describeGitCommitsInPageRequest := *openapiclient.NewDescribeGitCommitsInPageRequest(int64(123), "Ref_example") // DescribeGitCommitsInPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitCommitsInPage(context.Background()).Authorization(authorization).Action(action).DescribeGitCommitsInPageRequest(describeGitCommitsInPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitCommitsInPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitCommitsInPage`: DescribeGitCommitsInPage200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitCommitsInPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitCommitsInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitCommitsInPage&quot;]
 **describeGitCommitsInPageRequest** | [**DescribeGitCommitsInPageRequest**](DescribeGitCommitsInPageRequest.md) |  | 

### Return type

[**DescribeGitCommitsInPage200Response**](DescribeGitCommitsInPage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitContributors

> DescribeGitContributors200Response DescribeGitContributors(ctx).Authorization(authorization).Action(action).DescribeGitContributorsRequest(describeGitContributorsRequest).Execute()

仓库信息-查询git仓库的贡献者



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitContributors" // string | Action (default to "DescribeGitContributors")
	describeGitContributorsRequest := *openapiclient.NewDescribeGitContributorsRequest(int64(123), "Ref_example") // DescribeGitContributorsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitContributors(context.Background()).Authorization(authorization).Action(action).DescribeGitContributorsRequest(describeGitContributorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitContributors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitContributors`: DescribeGitContributors200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitContributors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitContributors&quot;]
 **describeGitContributorsRequest** | [**DescribeGitContributorsRequest**](DescribeGitContributorsRequest.md) |  | 

### Return type

[**DescribeGitContributors200Response**](DescribeGitContributors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitDepot

> DescribeGitDepot200Response DescribeGitDepot(ctx).Authorization(authorization).Action(action).DescribeGitDepotRequest(describeGitDepotRequest).Execute()

仓库信息-根据代码仓库id获取代码仓库信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitDepot" // string | Action (default to "DescribeGitDepot")
	describeGitDepotRequest := *openapiclient.NewDescribeGitDepotRequest(int64(123)) // DescribeGitDepotRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitDepot(context.Background()).Authorization(authorization).Action(action).DescribeGitDepotRequest(describeGitDepotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitDepot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitDepot`: DescribeGitDepot200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitDepot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitDepotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitDepot&quot;]
 **describeGitDepotRequest** | [**DescribeGitDepotRequest**](DescribeGitDepotRequest.md) |  | 

### Return type

[**DescribeGitDepot200Response**](DescribeGitDepot200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitDepotDeployKeys

> DescribeGitDepotDeployKeys200Response DescribeGitDepotDeployKeys(ctx).Authorization(authorization).Action(action).DescribeGitDepotDeployKeysRequest(describeGitDepotDeployKeysRequest).Execute()

仓库设置-查询某仓库下的部署公钥列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitDepotDeployKeys" // string | Action (default to "DescribeGitDepotDeployKeys")
	describeGitDepotDeployKeysRequest := *openapiclient.NewDescribeGitDepotDeployKeysRequest(int64(123)) // DescribeGitDepotDeployKeysRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitDepotDeployKeys(context.Background()).Authorization(authorization).Action(action).DescribeGitDepotDeployKeysRequest(describeGitDepotDeployKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitDepotDeployKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitDepotDeployKeys`: DescribeGitDepotDeployKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitDepotDeployKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitDepotDeployKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitDepotDeployKeys&quot;]
 **describeGitDepotDeployKeysRequest** | [**DescribeGitDepotDeployKeysRequest**](DescribeGitDepotDeployKeysRequest.md) |  | 

### Return type

[**DescribeGitDepotDeployKeys200Response**](DescribeGitDepotDeployKeys200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitFile

> DescribeGitFile200Response DescribeGitFile(ctx).Authorization(authorization).Action(action).DescribeGitFileRequest(describeGitFileRequest).Execute()

Git文件-获取文件详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitFile" // string | Action (default to "DescribeGitFile")
	describeGitFileRequest := *openapiclient.NewDescribeGitFileRequest(int64(123), "Ref_example") // DescribeGitFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitFile(context.Background()).Authorization(authorization).Action(action).DescribeGitFileRequest(describeGitFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitFile`: DescribeGitFile200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitFile&quot;]
 **describeGitFileRequest** | [**DescribeGitFileRequest**](DescribeGitFileRequest.md) |  | 

### Return type

[**DescribeGitFile200Response**](DescribeGitFile200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitFileContent

> DescribeGitFileContent200Response DescribeGitFileContent(ctx).Authorization(authorization).Action(action).DescribeGitFileContentRequest(describeGitFileContentRequest).Execute()

Git提交-查询某次提交某文件的内容



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitFileContent" // string | Action (default to "DescribeGitFileContent")
	describeGitFileContentRequest := *openapiclient.NewDescribeGitFileContentRequest("CommitSha_example", int64(123), "Path_example") // DescribeGitFileContentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitFileContent(context.Background()).Authorization(authorization).Action(action).DescribeGitFileContentRequest(describeGitFileContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitFileContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitFileContent`: DescribeGitFileContent200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitFileContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitFileContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitFileContent&quot;]
 **describeGitFileContentRequest** | [**DescribeGitFileContentRequest**](DescribeGitFileContentRequest.md) |  | 

### Return type

[**DescribeGitFileContent200Response**](DescribeGitFileContent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitFileStat

> DescribeGitFileStat200Response DescribeGitFileStat(ctx).Authorization(authorization).Action(action).DescribeGitFileStatRequest(describeGitFileStatRequest).Execute()

Git文件-检查仓库文件是否存在



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitFileStat" // string | Action (default to "DescribeGitFileStat")
	describeGitFileStatRequest := *openapiclient.NewDescribeGitFileStatRequest("DepotPath_example", "Path_example", "Ref_example") // DescribeGitFileStatRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitFileStat(context.Background()).Authorization(authorization).Action(action).DescribeGitFileStatRequest(describeGitFileStatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitFileStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitFileStat`: DescribeGitFileStat200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitFileStat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitFileStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitFileStat&quot;]
 **describeGitFileStatRequest** | [**DescribeGitFileStatRequest**](DescribeGitFileStatRequest.md) |  | 

### Return type

[**DescribeGitFileStat200Response**](DescribeGitFileStat200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitFiles

> DescribeGitFiles200Response DescribeGitFiles(ctx).Authorization(authorization).Action(action).DescribeGitFilesRequest(describeGitFilesRequest).Execute()

Git文件-查询仓库目录下文件和文件夹名字



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitFiles" // string | Action (default to "DescribeGitFiles")
	describeGitFilesRequest := *openapiclient.NewDescribeGitFilesRequest(int64(123), "Ref_example") // DescribeGitFilesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitFiles(context.Background()).Authorization(authorization).Action(action).DescribeGitFilesRequest(describeGitFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitFiles`: DescribeGitFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitFiles&quot;]
 **describeGitFilesRequest** | [**DescribeGitFilesRequest**](DescribeGitFilesRequest.md) |  | 

### Return type

[**DescribeGitFiles200Response**](DescribeGitFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitMergeBase

> DescribeGitMergeBase200Response DescribeGitMergeBase(ctx).Authorization(authorization).Action(action).DescribeGitMergeBaseRequest(describeGitMergeBaseRequest).Execute()

仓库分支-查询两个分支的公共祖先



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitMergeBase" // string | Action (default to "DescribeGitMergeBase")
	describeGitMergeBaseRequest := *openapiclient.NewDescribeGitMergeBaseRequest(int64(123), "DestRef_example", "SrcRef_example") // DescribeGitMergeBaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitMergeBase(context.Background()).Authorization(authorization).Action(action).DescribeGitMergeBaseRequest(describeGitMergeBaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitMergeBase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitMergeBase`: DescribeGitMergeBase200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitMergeBase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitMergeBaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitMergeBase&quot;]
 **describeGitMergeBaseRequest** | [**DescribeGitMergeBaseRequest**](DescribeGitMergeBaseRequest.md) |  | 

### Return type

[**DescribeGitMergeBase200Response**](DescribeGitMergeBase200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitMergeRequestDiffDetail

> DescribeGitMergeRequestDiffDetail200Response DescribeGitMergeRequestDiffDetail(ctx).Authorization(authorization).Action(action).DescribeGitMergeRequestDiffDetailRequest(describeGitMergeRequestDiffDetailRequest).Execute()

合并请求-查询合并请求文件的 diff 详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitMergeRequestDiffDetail" // string | Action (default to "DescribeGitMergeRequestDiffDetail")
	describeGitMergeRequestDiffDetailRequest := *openapiclient.NewDescribeGitMergeRequestDiffDetailRequest(int64(123), int64(123), "Path_example") // DescribeGitMergeRequestDiffDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitMergeRequestDiffDetail(context.Background()).Authorization(authorization).Action(action).DescribeGitMergeRequestDiffDetailRequest(describeGitMergeRequestDiffDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitMergeRequestDiffDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitMergeRequestDiffDetail`: DescribeGitMergeRequestDiffDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitMergeRequestDiffDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitMergeRequestDiffDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitMergeRequestDiffDetail&quot;]
 **describeGitMergeRequestDiffDetailRequest** | [**DescribeGitMergeRequestDiffDetailRequest**](DescribeGitMergeRequestDiffDetailRequest.md) |  | 

### Return type

[**DescribeGitMergeRequestDiffDetail200Response**](DescribeGitMergeRequestDiffDetail200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitMergeRequestDiffs

> DescribeGitMergeRequestDiffs200Response DescribeGitMergeRequestDiffs(ctx).Authorization(authorization).Action(action).DescribeGitMergeRequestDiffsRequest(describeGitMergeRequestDiffsRequest).Execute()

合并请求-查询合并请求diff信息的文件列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitMergeRequestDiffs" // string | Action (default to "DescribeGitMergeRequestDiffs")
	describeGitMergeRequestDiffsRequest := *openapiclient.NewDescribeGitMergeRequestDiffsRequest(int64(123), int64(123)) // DescribeGitMergeRequestDiffsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitMergeRequestDiffs(context.Background()).Authorization(authorization).Action(action).DescribeGitMergeRequestDiffsRequest(describeGitMergeRequestDiffsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitMergeRequestDiffs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitMergeRequestDiffs`: DescribeGitMergeRequestDiffs200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitMergeRequestDiffs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitMergeRequestDiffsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitMergeRequestDiffs&quot;]
 **describeGitMergeRequestDiffsRequest** | [**DescribeGitMergeRequestDiffsRequest**](DescribeGitMergeRequestDiffsRequest.md) |  | 

### Return type

[**DescribeGitMergeRequestDiffs200Response**](DescribeGitMergeRequestDiffs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitMergeRequestParticipants

> DescribeGitMergeRequestParticipants200Response DescribeGitMergeRequestParticipants(ctx).Authorization(authorization).Action(action).DescribeGitMergeRequestParticipantsRequest(describeGitMergeRequestParticipantsRequest).Execute()

合并请求-获取合并请求的参与者



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitMergeRequestParticipants" // string | Action (default to "DescribeGitMergeRequestParticipants")
	describeGitMergeRequestParticipantsRequest := *openapiclient.NewDescribeGitMergeRequestParticipantsRequest(int64(123), int64(123)) // DescribeGitMergeRequestParticipantsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitMergeRequestParticipants(context.Background()).Authorization(authorization).Action(action).DescribeGitMergeRequestParticipantsRequest(describeGitMergeRequestParticipantsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitMergeRequestParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitMergeRequestParticipants`: DescribeGitMergeRequestParticipants200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitMergeRequestParticipants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitMergeRequestParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitMergeRequestParticipants&quot;]
 **describeGitMergeRequestParticipantsRequest** | [**DescribeGitMergeRequestParticipantsRequest**](DescribeGitMergeRequestParticipantsRequest.md) |  | 

### Return type

[**DescribeGitMergeRequestParticipants200Response**](DescribeGitMergeRequestParticipants200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitMergeRequestsBySha

> DescribeGitMergeRequestsBySha200Response DescribeGitMergeRequestsBySha(ctx).Authorization(authorization).Action(action).DescribeGitMergeRequestsByShaRequest(describeGitMergeRequestsByShaRequest).Execute()

合并请求-查询含有某次提交的合并请求



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitMergeRequestsBySha" // string | Action (default to "DescribeGitMergeRequestsBySha")
	describeGitMergeRequestsByShaRequest := *openapiclient.NewDescribeGitMergeRequestsByShaRequest(int64(123), "Sha_example") // DescribeGitMergeRequestsByShaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitMergeRequestsBySha(context.Background()).Authorization(authorization).Action(action).DescribeGitMergeRequestsByShaRequest(describeGitMergeRequestsByShaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitMergeRequestsBySha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitMergeRequestsBySha`: DescribeGitMergeRequestsBySha200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitMergeRequestsBySha`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitMergeRequestsByShaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitMergeRequestsBySha&quot;]
 **describeGitMergeRequestsByShaRequest** | [**DescribeGitMergeRequestsByShaRequest**](DescribeGitMergeRequestsByShaRequest.md) |  | 

### Return type

[**DescribeGitMergeRequestsBySha200Response**](DescribeGitMergeRequestsBySha200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitProjectDeployKeys

> DescribeGitDepotDeployKeys200Response DescribeGitProjectDeployKeys(ctx).Authorization(authorization).Action(action).DescribeGitProjectDeployKeysRequest(describeGitProjectDeployKeysRequest).Execute()

仓库设置-查询某项目下的部署公钥列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitProjectDeployKeys" // string | Action (default to "DescribeGitProjectDeployKeys")
	describeGitProjectDeployKeysRequest := *openapiclient.NewDescribeGitProjectDeployKeysRequest(int64(123)) // DescribeGitProjectDeployKeysRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitProjectDeployKeys(context.Background()).Authorization(authorization).Action(action).DescribeGitProjectDeployKeysRequest(describeGitProjectDeployKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitProjectDeployKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitProjectDeployKeys`: DescribeGitDepotDeployKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitProjectDeployKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitProjectDeployKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitProjectDeployKeys&quot;]
 **describeGitProjectDeployKeysRequest** | [**DescribeGitProjectDeployKeysRequest**](DescribeGitProjectDeployKeysRequest.md) |  | 

### Return type

[**DescribeGitDepotDeployKeys200Response**](DescribeGitDepotDeployKeys200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitProtectedTags

> DescribeGitProtectedTags200Response DescribeGitProtectedTags(ctx).Authorization(authorization).Action(action).DescribeGitProtectedTagsRequest(describeGitProtectedTagsRequest).Execute()

标签信息-查询受保护的标签列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitProtectedTags" // string | Action (default to "DescribeGitProtectedTags")
	describeGitProtectedTagsRequest := *openapiclient.NewDescribeGitProtectedTagsRequest(int64(123)) // DescribeGitProtectedTagsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitProtectedTags(context.Background()).Authorization(authorization).Action(action).DescribeGitProtectedTagsRequest(describeGitProtectedTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitProtectedTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitProtectedTags`: DescribeGitProtectedTags200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitProtectedTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitProtectedTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitProtectedTags&quot;]
 **describeGitProtectedTagsRequest** | [**DescribeGitProtectedTagsRequest**](DescribeGitProtectedTagsRequest.md) |  | 

### Return type

[**DescribeGitProtectedTags200Response**](DescribeGitProtectedTags200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitProtectedTagsByRule

> DescribeGitProtectedTagsByRule200Response DescribeGitProtectedTagsByRule(ctx).Authorization(authorization).Action(action).DescribeGitProtectedTagsByRuleRequest(describeGitProtectedTagsByRuleRequest).Execute()

标签信息-根据标签保护规则查询受保护的标签列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitProtectedTagsByRule" // string | Action (default to "DescribeGitProtectedTagsByRule")
	describeGitProtectedTagsByRuleRequest := *openapiclient.NewDescribeGitProtectedTagsByRuleRequest(int64(123), "Rule_example") // DescribeGitProtectedTagsByRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitProtectedTagsByRule(context.Background()).Authorization(authorization).Action(action).DescribeGitProtectedTagsByRuleRequest(describeGitProtectedTagsByRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitProtectedTagsByRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitProtectedTagsByRule`: DescribeGitProtectedTagsByRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitProtectedTagsByRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitProtectedTagsByRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitProtectedTagsByRule&quot;]
 **describeGitProtectedTagsByRuleRequest** | [**DescribeGitProtectedTagsByRuleRequest**](DescribeGitProtectedTagsByRuleRequest.md) |  | 

### Return type

[**DescribeGitProtectedTagsByRule200Response**](DescribeGitProtectedTagsByRule200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitRef

> DescribeGitRef200Response DescribeGitRef(ctx).Authorization(authorization).Action(action).DescribeGitRefRequest(describeGitRefRequest).Execute()

仓库分支-获取分支或标签信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitRef" // string | Action (default to "DescribeGitRef")
	describeGitRefRequest := *openapiclient.NewDescribeGitRefRequest(int64(123), "Revision_example") // DescribeGitRefRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitRef(context.Background()).Authorization(authorization).Action(action).DescribeGitRefRequest(describeGitRefRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitRef`: DescribeGitRef200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitRef`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitRef&quot;]
 **describeGitRefRequest** | [**DescribeGitRefRequest**](DescribeGitRefRequest.md) |  | 

### Return type

[**DescribeGitRef200Response**](DescribeGitRef200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitRefsBySha

> DescribeGitRefsBySha200Response DescribeGitRefsBySha(ctx).Authorization(authorization).Action(action).DescribeGitRefsByShaRequest(describeGitRefsByShaRequest).Execute()

Git提交-查询含有某次提交的标签或分支列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitRefsBySha" // string | Action (default to "DescribeGitRefsBySha")
	describeGitRefsByShaRequest := *openapiclient.NewDescribeGitRefsByShaRequest(int64(123), "Sha_example", "Type_example") // DescribeGitRefsByShaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitRefsBySha(context.Background()).Authorization(authorization).Action(action).DescribeGitRefsByShaRequest(describeGitRefsByShaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitRefsBySha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitRefsBySha`: DescribeGitRefsBySha200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitRefsBySha`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitRefsByShaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitRefsBySha&quot;]
 **describeGitRefsByShaRequest** | [**DescribeGitRefsByShaRequest**](DescribeGitRefsByShaRequest.md) |  | 

### Return type

[**DescribeGitRefsBySha200Response**](DescribeGitRefsBySha200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitReleaseDetail

> DescribeGitReleaseDetail200Response DescribeGitReleaseDetail(ctx).Authorization(authorization).Action(action).DescribeGitReleaseDetailRequest(describeGitReleaseDetailRequest).Execute()

版本信息-查询仓库的版本详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitReleaseDetail" // string | Action (default to "DescribeGitReleaseDetail")
	describeGitReleaseDetailRequest := *openapiclient.NewDescribeGitReleaseDetailRequest(int64(123), int64(123), "TagName_example") // DescribeGitReleaseDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitReleaseDetail(context.Background()).Authorization(authorization).Action(action).DescribeGitReleaseDetailRequest(describeGitReleaseDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitReleaseDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitReleaseDetail`: DescribeGitReleaseDetail200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitReleaseDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitReleaseDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitReleaseDetail&quot;]
 **describeGitReleaseDetailRequest** | [**DescribeGitReleaseDetailRequest**](DescribeGitReleaseDetailRequest.md) |  | 

### Return type

[**DescribeGitReleaseDetail200Response**](DescribeGitReleaseDetail200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitReleases

> DescribeGitReleases200Response DescribeGitReleases(ctx).Authorization(authorization).Action(action).DescribeGitReleasesRequest(describeGitReleasesRequest).Execute()

版本信息-查询仓库的版本列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitReleases" // string | Action (default to "DescribeGitReleases")
	describeGitReleasesRequest := *openapiclient.NewDescribeGitReleasesRequest(int64(123)) // DescribeGitReleasesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitReleases(context.Background()).Authorization(authorization).Action(action).DescribeGitReleasesRequest(describeGitReleasesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitReleases`: DescribeGitReleases200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitReleases&quot;]
 **describeGitReleasesRequest** | [**DescribeGitReleasesRequest**](DescribeGitReleasesRequest.md) |  | 

### Return type

[**DescribeGitReleases200Response**](DescribeGitReleases200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitTag

> DescribeGitTag200Response DescribeGitTag(ctx).Authorization(authorization).Action(action).DescribeGitTagRequest(describeGitTagRequest).Execute()

标签信息-查询单个tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitTag" // string | Action (default to "DescribeGitTag")
	describeGitTagRequest := *openapiclient.NewDescribeGitTagRequest(int64(123), "TagName_example") // DescribeGitTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitTag(context.Background()).Authorization(authorization).Action(action).DescribeGitTagRequest(describeGitTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitTag`: DescribeGitTag200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitTag&quot;]
 **describeGitTagRequest** | [**DescribeGitTagRequest**](DescribeGitTagRequest.md) |  | 

### Return type

[**DescribeGitTag200Response**](DescribeGitTag200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitTags

> DescribeGitTags200Response DescribeGitTags(ctx).Authorization(authorization).Action(action).DescribeGitTagsRequest(describeGitTagsRequest).Execute()

标签信息-查询当前仓库下所有tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitTags" // string | Action (default to "DescribeGitTags")
	describeGitTagsRequest := *openapiclient.NewDescribeGitTagsRequest(int64(123)) // DescribeGitTagsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitTags(context.Background()).Authorization(authorization).Action(action).DescribeGitTagsRequest(describeGitTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitTags`: DescribeGitTags200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitTags&quot;]
 **describeGitTagsRequest** | [**DescribeGitTagsRequest**](DescribeGitTagsRequest.md) |  | 

### Return type

[**DescribeGitTags200Response**](DescribeGitTags200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitTagsByBranch

> DescribeGitTagsByBranch200Response DescribeGitTagsByBranch(ctx).Authorization(authorization).Action(action).DescribeGitTagsByBranchRequest(describeGitTagsByBranchRequest).Execute()

标签信息-根据分支获取标签列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitTagsByBranch" // string | Action (default to "DescribeGitTagsByBranch")
	describeGitTagsByBranchRequest := *openapiclient.NewDescribeGitTagsByBranchRequest("Branch_example", int64(123)) // DescribeGitTagsByBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitTagsByBranch(context.Background()).Authorization(authorization).Action(action).DescribeGitTagsByBranchRequest(describeGitTagsByBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitTagsByBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitTagsByBranch`: DescribeGitTagsByBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitTagsByBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitTagsByBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitTagsByBranch&quot;]
 **describeGitTagsByBranchRequest** | [**DescribeGitTagsByBranchRequest**](DescribeGitTagsByBranchRequest.md) |  | 

### Return type

[**DescribeGitTagsByBranch200Response**](DescribeGitTagsByBranch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitTagsBySha

> DescribeGitTagsBySha200Response DescribeGitTagsBySha(ctx).Authorization(authorization).Action(action).DescribeGitTagsByShaRequest(describeGitTagsByShaRequest).Execute()

标签信息-查询含有某次提交的标签列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitTagsBySha" // string | Action (default to "DescribeGitTagsBySha")
	describeGitTagsByShaRequest := *openapiclient.NewDescribeGitTagsByShaRequest(int64(123), "Sha_example") // DescribeGitTagsByShaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitTagsBySha(context.Background()).Authorization(authorization).Action(action).DescribeGitTagsByShaRequest(describeGitTagsByShaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitTagsBySha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitTagsBySha`: DescribeGitTagsBySha200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitTagsBySha`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitTagsByShaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitTagsBySha&quot;]
 **describeGitTagsByShaRequest** | [**DescribeGitTagsByShaRequest**](DescribeGitTagsByShaRequest.md) |  | 

### Return type

[**DescribeGitTagsBySha200Response**](DescribeGitTagsBySha200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGitTree

> DescribeGitTree200Response DescribeGitTree(ctx).Authorization(authorization).Action(action).DescribeGitTreeRequest(describeGitTreeRequest).Execute()

仓库信息-查询git仓库的树



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGitTree" // string | Action (default to "DescribeGitTree")
	describeGitTreeRequest := *openapiclient.NewDescribeGitTreeRequest(int64(123), false, "Ref_example") // DescribeGitTreeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGitTree(context.Background()).Authorization(authorization).Action(action).DescribeGitTreeRequest(describeGitTreeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGitTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGitTree`: DescribeGitTree200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGitTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGitTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGitTree&quot;]
 **describeGitTreeRequest** | [**DescribeGitTreeRequest**](DescribeGitTreeRequest.md) |  | 

### Return type

[**DescribeGitTree200Response**](DescribeGitTree200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGrantObjectsOnResource

> DescribeGrantObjectsOnResource200Response DescribeGrantObjectsOnResource(ctx).Authorization(authorization).Action(action).DescribeGrantObjectsOnResourceRequest(describeGrantObjectsOnResourceRequest).Execute()

授权对象列表分页查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGrantObjectsOnResource" // string | Action (default to "DescribeGrantObjectsOnResource")
	describeGrantObjectsOnResourceRequest := *openapiclient.NewDescribeGrantObjectsOnResourceRequest(int32(123), int32(123), *openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example")) // DescribeGrantObjectsOnResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGrantObjectsOnResource(context.Background()).Authorization(authorization).Action(action).DescribeGrantObjectsOnResourceRequest(describeGrantObjectsOnResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGrantObjectsOnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGrantObjectsOnResource`: DescribeGrantObjectsOnResource200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGrantObjectsOnResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGrantObjectsOnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGrantObjectsOnResource&quot;]
 **describeGrantObjectsOnResourceRequest** | [**DescribeGrantObjectsOnResourceRequest**](DescribeGrantObjectsOnResourceRequest.md) |  | 

### Return type

[**DescribeGrantObjectsOnResource200Response**](DescribeGrantObjectsOnResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeGrantUsersOnResource

> DescribeGrantUsersOnResource200Response DescribeGrantUsersOnResource(ctx).Authorization(authorization).Action(action).DescribeGrantUsersOnResourceRequest(describeGrantUsersOnResourceRequest).Execute()

授权用户列表分页查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeGrantUsersOnResource" // string | Action (default to "DescribeGrantUsersOnResource")
	describeGrantUsersOnResourceRequest := *openapiclient.NewDescribeGrantUsersOnResourceRequest(int64(123), int64(123), *openapiclient.NewRamGrantResourceInfoRequest("ResourceId_example", "ResourceType_example")) // DescribeGrantUsersOnResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeGrantUsersOnResource(context.Background()).Authorization(authorization).Action(action).DescribeGrantUsersOnResourceRequest(describeGrantUsersOnResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeGrantUsersOnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGrantUsersOnResource`: DescribeGrantUsersOnResource200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeGrantUsersOnResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGrantUsersOnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeGrantUsersOnResource&quot;]
 **describeGrantUsersOnResourceRequest** | [**DescribeGrantUsersOnResourceRequest**](DescribeGrantUsersOnResourceRequest.md) |  | 

### Return type

[**DescribeGrantUsersOnResource200Response**](DescribeGrantUsersOnResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeHostServerInstance

> DescribeHostServerInstance200Response DescribeHostServerInstance(ctx).Authorization(authorization).Action(action).DescribeHostServerInstanceRequest(describeHostServerInstanceRequest).Execute()

CD 主机组部署详情获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeHostServerInstance" // string | Action (default to "DescribeHostServerInstance")
	describeHostServerInstanceRequest := *openapiclient.NewDescribeHostServerInstanceRequest("Account_example", "ServerGroupName_example") // DescribeHostServerInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeHostServerInstance(context.Background()).Authorization(authorization).Action(action).DescribeHostServerInstanceRequest(describeHostServerInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeHostServerInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeHostServerInstance`: DescribeHostServerInstance200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeHostServerInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeHostServerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeHostServerInstance&quot;]
 **describeHostServerInstanceRequest** | [**DescribeHostServerInstanceRequest**](DescribeHostServerInstanceRequest.md) |  | 

### Return type

[**DescribeHostServerInstance200Response**](DescribeHostServerInstance200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssue

> DescribeIssue200Response DescribeIssue(ctx).Authorization(authorization).Action(action).DescribeIssueRequest(describeIssueRequest).Execute()

事项详情查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssue" // string | Action (default to "DescribeIssue")
	describeIssueRequest := *openapiclient.NewDescribeIssueRequest(int64(123), "ProjectName_example") // DescribeIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssue(context.Background()).Authorization(authorization).Action(action).DescribeIssueRequest(describeIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssue`: DescribeIssue200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssue&quot;]
 **describeIssueRequest** | [**DescribeIssueRequest**](DescribeIssueRequest.md) |  | 

### Return type

[**DescribeIssue200Response**](DescribeIssue200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueAttachmentPreSignedUrl

> DescribeIssueAttachmentPreSignedUrl200Response DescribeIssueAttachmentPreSignedUrl(ctx).Authorization(authorization).Action(action).DescribeIssueAttachmentPreSignedUrlRequest(describeIssueAttachmentPreSignedUrlRequest).Execute()

预签名信息获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueAttachmentPreSignedUrl" // string | Action (default to "DescribeIssueAttachmentPreSignedUrl")
	describeIssueAttachmentPreSignedUrlRequest := *openapiclient.NewDescribeIssueAttachmentPreSignedUrlRequest("FileName_example", int64(123), "ProjectName_example") // DescribeIssueAttachmentPreSignedUrlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueAttachmentPreSignedUrl(context.Background()).Authorization(authorization).Action(action).DescribeIssueAttachmentPreSignedUrlRequest(describeIssueAttachmentPreSignedUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueAttachmentPreSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueAttachmentPreSignedUrl`: DescribeIssueAttachmentPreSignedUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueAttachmentPreSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueAttachmentPreSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueAttachmentPreSignedUrl&quot;]
 **describeIssueAttachmentPreSignedUrlRequest** | [**DescribeIssueAttachmentPreSignedUrlRequest**](DescribeIssueAttachmentPreSignedUrlRequest.md) |  | 

### Return type

[**DescribeIssueAttachmentPreSignedUrl200Response**](DescribeIssueAttachmentPreSignedUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueCommentList

> DescribeIssueCommentList200Response DescribeIssueCommentList(ctx).Authorization(authorization).Action(action).DescribeIssueCommentListRequest(describeIssueCommentListRequest).Execute()

事项评论列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueCommentList" // string | Action (default to "DescribeIssueCommentList")
	describeIssueCommentListRequest := *openapiclient.NewDescribeIssueCommentListRequest(int64(123), "ProjectName_example") // DescribeIssueCommentListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueCommentList(context.Background()).Authorization(authorization).Action(action).DescribeIssueCommentListRequest(describeIssueCommentListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueCommentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueCommentList`: DescribeIssueCommentList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueCommentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueCommentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueCommentList&quot;]
 **describeIssueCommentListRequest** | [**DescribeIssueCommentListRequest**](DescribeIssueCommentListRequest.md) |  | 

### Return type

[**DescribeIssueCommentList200Response**](DescribeIssueCommentList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueCustomFieldLogList

> DescribeIssueCustomFieldLogList200Response DescribeIssueCustomFieldLogList(ctx).Authorization(authorization).Action(action).DescribeIssueCustomFieldLogListRequest(describeIssueCustomFieldLogListRequest).Execute()

事项的自定义属性变更日志查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueCustomFieldLogList" // string | Action (default to "DescribeIssueCustomFieldLogList")
	describeIssueCustomFieldLogListRequest := *openapiclient.NewDescribeIssueCustomFieldLogListRequest("FieldName_example", int64(123), "ProjectName_example") // DescribeIssueCustomFieldLogListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueCustomFieldLogList(context.Background()).Authorization(authorization).Action(action).DescribeIssueCustomFieldLogListRequest(describeIssueCustomFieldLogListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueCustomFieldLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueCustomFieldLogList`: DescribeIssueCustomFieldLogList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueCustomFieldLogList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueCustomFieldLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueCustomFieldLogList&quot;]
 **describeIssueCustomFieldLogListRequest** | [**DescribeIssueCustomFieldLogListRequest**](DescribeIssueCustomFieldLogListRequest.md) |  | 

### Return type

[**DescribeIssueCustomFieldLogList200Response**](DescribeIssueCustomFieldLogList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueFileUrl

> DescribeIssueFileUrl200Response DescribeIssueFileUrl(ctx).Authorization(authorization).Action(action).DescribeIssueFileUrlRequest(describeIssueFileUrlRequest).Execute()

事项附件的下载地址查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueFileUrl" // string | Action (default to "DescribeIssueFileUrl")
	describeIssueFileUrlRequest := *openapiclient.NewDescribeIssueFileUrlRequest(int64(123), "ProjectName_example") // DescribeIssueFileUrlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueFileUrl(context.Background()).Authorization(authorization).Action(action).DescribeIssueFileUrlRequest(describeIssueFileUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueFileUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueFileUrl`: DescribeIssueFileUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueFileUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueFileUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueFileUrl&quot;]
 **describeIssueFileUrlRequest** | [**DescribeIssueFileUrlRequest**](DescribeIssueFileUrlRequest.md) |  | 

### Return type

[**DescribeIssueFileUrl200Response**](DescribeIssueFileUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueFilterList

> DescribeIssueFilterList200Response DescribeIssueFilterList(ctx).Authorization(authorization).Action(action).DescribeIssueFilterListRequest(describeIssueFilterListRequest).Execute()

事项筛选器列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueFilterList" // string | Action (default to "DescribeIssueFilterList")
	describeIssueFilterListRequest := *openapiclient.NewDescribeIssueFilterListRequest("IssueType_example", "ProjectName_example") // DescribeIssueFilterListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueFilterList(context.Background()).Authorization(authorization).Action(action).DescribeIssueFilterListRequest(describeIssueFilterListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueFilterList`: DescribeIssueFilterList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueFilterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueFilterList&quot;]
 **describeIssueFilterListRequest** | [**DescribeIssueFilterListRequest**](DescribeIssueFilterListRequest.md) |  | 

### Return type

[**DescribeIssueFilterList200Response**](DescribeIssueFilterList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueList

> DescribeIssueList200Response DescribeIssueList(ctx).Authorization(authorization).Action(action).DescribeIssueListRequest(describeIssueListRequest).Execute()

事项列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueList" // string | Action (default to "DescribeIssueList")
	describeIssueListRequest := *openapiclient.NewDescribeIssueListRequest("IssueType_example", "ProjectName_example") // DescribeIssueListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueList(context.Background()).Authorization(authorization).Action(action).DescribeIssueListRequest(describeIssueListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueList`: DescribeIssueList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueList&quot;]
 **describeIssueListRequest** | [**DescribeIssueListRequest**](DescribeIssueListRequest.md) |  | 

### Return type

[**DescribeIssueList200Response**](DescribeIssueList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueListWithPage

> DescribeIssueListWithPage200Response DescribeIssueListWithPage(ctx).Authorization(authorization).Action(action).DescribeIssueListWithPageRequest(describeIssueListWithPageRequest).Execute()

事项查询,返回分页信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueListWithPage" // string | Action (default to "DescribeIssueListWithPage")
	describeIssueListWithPageRequest := *openapiclient.NewDescribeIssueListWithPageRequest("IssueType_example", int64(123), int64(123), "ProjectName_example") // DescribeIssueListWithPageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueListWithPage(context.Background()).Authorization(authorization).Action(action).DescribeIssueListWithPageRequest(describeIssueListWithPageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueListWithPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueListWithPage`: DescribeIssueListWithPage200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueListWithPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueListWithPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueListWithPage&quot;]
 **describeIssueListWithPageRequest** | [**DescribeIssueListWithPageRequest**](DescribeIssueListWithPageRequest.md) |  | 

### Return type

[**DescribeIssueListWithPage200Response**](DescribeIssueListWithPage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueModuleList

> DescribeIssueModuleList200Response DescribeIssueModuleList(ctx).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()

事项模块列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueModuleList" // string | Action (default to "DescribeIssueModuleList")
	describeAPIDocListRequest := *openapiclient.NewDescribeAPIDocListRequest("ProjectName_example") // DescribeAPIDocListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueModuleList(context.Background()).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueModuleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueModuleList`: DescribeIssueModuleList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueModuleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueModuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueModuleList&quot;]
 **describeAPIDocListRequest** | [**DescribeAPIDocListRequest**](DescribeAPIDocListRequest.md) |  | 

### Return type

[**DescribeIssueModuleList200Response**](DescribeIssueModuleList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueReferenceResources

> DescribeIssueReferenceResources200Response DescribeIssueReferenceResources(ctx).Authorization(authorization).Action(action).DescribeIssueReferenceResourcesRequest(describeIssueReferenceResourcesRequest).Execute()

事项的引用资源列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueReferenceResources" // string | Action (default to "DescribeIssueReferenceResources")
	describeIssueReferenceResourcesRequest := *openapiclient.NewDescribeIssueReferenceResourcesRequest(int64(123), int64(123), "ProjectName_example") // DescribeIssueReferenceResourcesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueReferenceResources(context.Background()).Authorization(authorization).Action(action).DescribeIssueReferenceResourcesRequest(describeIssueReferenceResourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueReferenceResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueReferenceResources`: DescribeIssueReferenceResources200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueReferenceResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueReferenceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueReferenceResources&quot;]
 **describeIssueReferenceResourcesRequest** | [**DescribeIssueReferenceResourcesRequest**](DescribeIssueReferenceResourcesRequest.md) |  | 

### Return type

[**DescribeIssueReferenceResources200Response**](DescribeIssueReferenceResources200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueRelatedWorkItemList

> DescribeIssueRelatedWorkItemList200Response DescribeIssueRelatedWorkItemList(ctx).Authorization(authorization).Action(action).DescribeIssueRelatedWorkItemListRequest(describeIssueRelatedWorkItemListRequest).Execute()

事项关联的项目集中的工作项查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueRelatedWorkItemList" // string | Action (default to "DescribeIssueRelatedWorkItemList")
	describeIssueRelatedWorkItemListRequest := *openapiclient.NewDescribeIssueRelatedWorkItemListRequest(int64(123), "ProjectName_example") // DescribeIssueRelatedWorkItemListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueRelatedWorkItemList(context.Background()).Authorization(authorization).Action(action).DescribeIssueRelatedWorkItemListRequest(describeIssueRelatedWorkItemListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueRelatedWorkItemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueRelatedWorkItemList`: DescribeIssueRelatedWorkItemList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueRelatedWorkItemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueRelatedWorkItemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueRelatedWorkItemList&quot;]
 **describeIssueRelatedWorkItemListRequest** | [**DescribeIssueRelatedWorkItemListRequest**](DescribeIssueRelatedWorkItemListRequest.md) |  | 

### Return type

[**DescribeIssueRelatedWorkItemList200Response**](DescribeIssueRelatedWorkItemList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueReleaseList

> DescribeIssueReleaseList200Response DescribeIssueReleaseList(ctx).Authorization(authorization).Action(action).DescribeIssueReleaseListRequest(describeIssueReleaseListRequest).Execute()

事项加入的版本查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueReleaseList" // string | Action (default to "DescribeIssueReleaseList")
	describeIssueReleaseListRequest := *openapiclient.NewDescribeIssueReleaseListRequest(int64(123), "ProjectName_example") // DescribeIssueReleaseListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueReleaseList(context.Background()).Authorization(authorization).Action(action).DescribeIssueReleaseListRequest(describeIssueReleaseListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueReleaseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueReleaseList`: DescribeIssueReleaseList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueReleaseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueReleaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueReleaseList&quot;]
 **describeIssueReleaseListRequest** | [**DescribeIssueReleaseListRequest**](DescribeIssueReleaseListRequest.md) |  | 

### Return type

[**DescribeIssueReleaseList200Response**](DescribeIssueReleaseList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueStatusChangeLogList

> DescribeIssueStatusChangeLogList200Response DescribeIssueStatusChangeLogList(ctx).Authorization(authorization).Action(action).DescribeIssueStatusChangeLogListRequest(describeIssueStatusChangeLogListRequest).Execute()

事项的状态变更记录查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueStatusChangeLogList" // string | Action (default to "DescribeIssueStatusChangeLogList")
	describeIssueStatusChangeLogListRequest := *openapiclient.NewDescribeIssueStatusChangeLogListRequest([]int64{int64(123)}, "ProjectName_example") // DescribeIssueStatusChangeLogListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueStatusChangeLogList(context.Background()).Authorization(authorization).Action(action).DescribeIssueStatusChangeLogListRequest(describeIssueStatusChangeLogListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueStatusChangeLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueStatusChangeLogList`: DescribeIssueStatusChangeLogList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueStatusChangeLogList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueStatusChangeLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueStatusChangeLogList&quot;]
 **describeIssueStatusChangeLogListRequest** | [**DescribeIssueStatusChangeLogListRequest**](DescribeIssueStatusChangeLogListRequest.md) |  | 

### Return type

[**DescribeIssueStatusChangeLogList200Response**](DescribeIssueStatusChangeLogList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIssueWorkLogList

> DescribeIssueWorkLogList200Response DescribeIssueWorkLogList(ctx).Authorization(authorization).Action(action).DescribeIssueWorkLogListRequest(describeIssueWorkLogListRequest).Execute()

事项的工时日志查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIssueWorkLogList" // string | Action (default to "DescribeIssueWorkLogList")
	describeIssueWorkLogListRequest := *openapiclient.NewDescribeIssueWorkLogListRequest(int64(123), "ProjectName_example") // DescribeIssueWorkLogListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIssueWorkLogList(context.Background()).Authorization(authorization).Action(action).DescribeIssueWorkLogListRequest(describeIssueWorkLogListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIssueWorkLogList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIssueWorkLogList`: DescribeIssueWorkLogList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIssueWorkLogList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIssueWorkLogListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIssueWorkLogList&quot;]
 **describeIssueWorkLogListRequest** | [**DescribeIssueWorkLogListRequest**](DescribeIssueWorkLogListRequest.md) |  | 

### Return type

[**DescribeIssueWorkLogList200Response**](DescribeIssueWorkLogList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIteration

> DescribeIteration200Response DescribeIteration(ctx).Authorization(authorization).Action(action).DeleteIterationRequest(deleteIterationRequest).Execute()

迭代详情查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIteration" // string | Action (default to "DescribeIteration")
	deleteIterationRequest := *openapiclient.NewDeleteIterationRequest(int64(123), "ProjectName_example") // DeleteIterationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIteration(context.Background()).Authorization(authorization).Action(action).DeleteIterationRequest(deleteIterationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIteration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIteration`: DescribeIteration200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIteration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIteration&quot;]
 **deleteIterationRequest** | [**DeleteIterationRequest**](DeleteIterationRequest.md) |  | 

### Return type

[**DescribeIteration200Response**](DescribeIteration200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeIterationList

> DescribeIterationList200Response DescribeIterationList(ctx).Authorization(authorization).Action(action).DescribeIterationListRequest(describeIterationListRequest).Execute()

迭代列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeIterationList" // string | Action (default to "DescribeIterationList")
	describeIterationListRequest := *openapiclient.NewDescribeIterationListRequest("ProjectName_example") // DescribeIterationListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeIterationList(context.Background()).Authorization(authorization).Action(action).DescribeIterationListRequest(describeIterationListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeIterationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeIterationList`: DescribeIterationList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeIterationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeIterationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeIterationList&quot;]
 **describeIterationListRequest** | [**DescribeIterationListRequest**](DescribeIterationListRequest.md) |  | 

### Return type

[**DescribeIterationList200Response**](DescribeIterationList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMemberSshKey

> DescribeMemberSshKey200Response DescribeMemberSshKey(ctx).Authorization(authorization).Action(action).DescribeMemberSshKeyRequest(describeMemberSshKeyRequest).Execute()

仓库设置-获取团队成员的SSH公钥列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMemberSshKey" // string | Action (default to "DescribeMemberSshKey")
	describeMemberSshKeyRequest := *openapiclient.NewDescribeMemberSshKeyRequest(int64(123)) // DescribeMemberSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMemberSshKey(context.Background()).Authorization(authorization).Action(action).DescribeMemberSshKeyRequest(describeMemberSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMemberSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMemberSshKey`: DescribeMemberSshKey200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMemberSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMemberSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMemberSshKey&quot;]
 **describeMemberSshKeyRequest** | [**DescribeMemberSshKeyRequest**](DescribeMemberSshKeyRequest.md) |  | 

### Return type

[**DescribeMemberSshKey200Response**](DescribeMemberSshKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMergeReqCommits

> DescribeMergeReqCommits200Response DescribeMergeReqCommits(ctx).Authorization(authorization).Action(action).DescribeMergeReqCommitsRequest(describeMergeReqCommitsRequest).Execute()

合并请求-查询合并请求列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMergeReqCommits" // string | Action (default to "DescribeMergeReqCommits")
	describeMergeReqCommitsRequest := *openapiclient.NewDescribeMergeReqCommitsRequest(int64(123), int64(123)) // DescribeMergeReqCommitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMergeReqCommits(context.Background()).Authorization(authorization).Action(action).DescribeMergeReqCommitsRequest(describeMergeReqCommitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMergeReqCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMergeReqCommits`: DescribeMergeReqCommits200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMergeReqCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMergeReqCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMergeReqCommits&quot;]
 **describeMergeReqCommitsRequest** | [**DescribeMergeReqCommitsRequest**](DescribeMergeReqCommitsRequest.md) |  | 

### Return type

[**DescribeMergeReqCommits200Response**](DescribeMergeReqCommits200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMergeReqInfo

> DescribeMergeReqInfo200Response DescribeMergeReqInfo(ctx).Authorization(authorization).Action(action).DescribeMergeReqInfoRequest(describeMergeReqInfoRequest).Execute()

合并请求-查询合并请求详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMergeReqInfo" // string | Action (default to "DescribeMergeReqInfo")
	describeMergeReqInfoRequest := *openapiclient.NewDescribeMergeReqInfoRequest(int64(123), int64(123)) // DescribeMergeReqInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMergeReqInfo(context.Background()).Authorization(authorization).Action(action).DescribeMergeReqInfoRequest(describeMergeReqInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMergeReqInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMergeReqInfo`: DescribeMergeReqInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMergeReqInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMergeReqInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMergeReqInfo&quot;]
 **describeMergeReqInfoRequest** | [**DescribeMergeReqInfoRequest**](DescribeMergeReqInfoRequest.md) |  | 

### Return type

[**DescribeMergeReqInfo200Response**](DescribeMergeReqInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMergeRequest

> DescribeMergeRequest200Response DescribeMergeRequest(ctx).Authorization(authorization).Action(action).DescribeMergeRequestRequest(describeMergeRequestRequest).Execute()

合并请求-查询合并请求详情信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMergeRequest" // string | Action (default to "DescribeMergeRequest")
	describeMergeRequestRequest := *openapiclient.NewDescribeMergeRequestRequest(int64(123), int64(123)) // DescribeMergeRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMergeRequest(context.Background()).Authorization(authorization).Action(action).DescribeMergeRequestRequest(describeMergeRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMergeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMergeRequest`: DescribeMergeRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMergeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMergeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMergeRequest&quot;]
 **describeMergeRequestRequest** | [**DescribeMergeRequestRequest**](DescribeMergeRequestRequest.md) |  | 

### Return type

[**DescribeMergeRequest200Response**](DescribeMergeRequest200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMergeRequestFileDiff

> DescribeMergeRequestFileDiff200Response DescribeMergeRequestFileDiff(ctx).Authorization(authorization).Action(action).DescribeMergeRequestFileDiffRequest(describeMergeRequestFileDiffRequest).Execute()

合并请求-获取合并请求文件修改记录



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMergeRequestFileDiff" // string | Action (default to "DescribeMergeRequestFileDiff")
	describeMergeRequestFileDiffRequest := *openapiclient.NewDescribeMergeRequestFileDiffRequest(int64(123), int64(123)) // DescribeMergeRequestFileDiffRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMergeRequestFileDiff(context.Background()).Authorization(authorization).Action(action).DescribeMergeRequestFileDiffRequest(describeMergeRequestFileDiffRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMergeRequestFileDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMergeRequestFileDiff`: DescribeMergeRequestFileDiff200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMergeRequestFileDiff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMergeRequestFileDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMergeRequestFileDiff&quot;]
 **describeMergeRequestFileDiffRequest** | [**DescribeMergeRequestFileDiffRequest**](DescribeMergeRequestFileDiffRequest.md) |  | 

### Return type

[**DescribeMergeRequestFileDiff200Response**](DescribeMergeRequestFileDiff200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMergeRequestLog

> DescribeMergeRequestLog200Response DescribeMergeRequestLog(ctx).Authorization(authorization).Action(action).ModifyCloseMRRequest(modifyCloseMRRequest).Execute()

合并请求-查询合并请求操作记录



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMergeRequestLog" // string | Action (default to "DescribeMergeRequestLog")
	modifyCloseMRRequest := *openapiclient.NewModifyCloseMRRequest(int64(123), int64(123)) // ModifyCloseMRRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMergeRequestLog(context.Background()).Authorization(authorization).Action(action).ModifyCloseMRRequest(modifyCloseMRRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMergeRequestLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMergeRequestLog`: DescribeMergeRequestLog200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMergeRequestLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMergeRequestLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMergeRequestLog&quot;]
 **modifyCloseMRRequest** | [**ModifyCloseMRRequest**](ModifyCloseMRRequest.md) |  | 

### Return type

[**DescribeMergeRequestLog200Response**](DescribeMergeRequestLog200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMergeRequestReviewers

> DescribeMergeRequestReviewers200Response DescribeMergeRequestReviewers(ctx).Authorization(authorization).Action(action).DescribeMergeRequestReviewersRequest(describeMergeRequestReviewersRequest).Execute()

合并请求-获取合并请求的评审者



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMergeRequestReviewers" // string | Action (default to "DescribeMergeRequestReviewers")
	describeMergeRequestReviewersRequest := *openapiclient.NewDescribeMergeRequestReviewersRequest(int64(123), int64(123)) // DescribeMergeRequestReviewersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMergeRequestReviewers(context.Background()).Authorization(authorization).Action(action).DescribeMergeRequestReviewersRequest(describeMergeRequestReviewersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMergeRequestReviewers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMergeRequestReviewers`: DescribeMergeRequestReviewers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMergeRequestReviewers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMergeRequestReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMergeRequestReviewers&quot;]
 **describeMergeRequestReviewersRequest** | [**DescribeMergeRequestReviewersRequest**](DescribeMergeRequestReviewersRequest.md) |  | 

### Return type

[**DescribeMergeRequestReviewers200Response**](DescribeMergeRequestReviewers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeMyDepots

> DescribeMyDepots200Response DescribeMyDepots(ctx).Authorization(authorization).Action(action).DescribeMyDepotsRequest(describeMyDepotsRequest).Execute()

仓库信息-获取当前用户拥有读权限的仓库列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeMyDepots" // string | Action (default to "DescribeMyDepots")
	describeMyDepotsRequest := *openapiclient.NewDescribeMyDepotsRequest(int64(123), int64(123)) // DescribeMyDepotsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeMyDepots(context.Background()).Authorization(authorization).Action(action).DescribeMyDepotsRequest(describeMyDepotsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeMyDepots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeMyDepots`: DescribeMyDepots200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeMyDepots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeMyDepotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeMyDepots&quot;]
 **describeMyDepotsRequest** | [**DescribeMyDepotsRequest**](DescribeMyDepotsRequest.md) |  | 

### Return type

[**DescribeMyDepots200Response**](DescribeMyDepots200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeNotesByCommits

> DescribeNotesByCommits200Response DescribeNotesByCommits(ctx).Authorization(authorization).Action(action).DescribeNotesByCommitsRequest(describeNotesByCommitsRequest).Execute()

仓库信息-获取提交的note信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeNotesByCommits" // string | Action (default to "DescribeNotesByCommits")
	describeNotesByCommitsRequest := *openapiclient.NewDescribeNotesByCommitsRequest([]string{"Commits_example"}, int64(123), "NoteRef_example") // DescribeNotesByCommitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeNotesByCommits(context.Background()).Authorization(authorization).Action(action).DescribeNotesByCommitsRequest(describeNotesByCommitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeNotesByCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeNotesByCommits`: DescribeNotesByCommits200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeNotesByCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNotesByCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeNotesByCommits&quot;]
 **describeNotesByCommitsRequest** | [**DescribeNotesByCommitsRequest**](DescribeNotesByCommitsRequest.md) |  | 

### Return type

[**DescribeNotesByCommits200Response**](DescribeNotesByCommits200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeOneProject

> DescribeOneProject200Response DescribeOneProject(ctx).Authorization(authorization).Action(action).DescribeOneProjectRequest(describeOneProjectRequest).Execute()

单个项目查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeOneProject" // string | Action (default to "DescribeOneProject")
	describeOneProjectRequest := *openapiclient.NewDescribeOneProjectRequest(int32(123)) // DescribeOneProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeOneProject(context.Background()).Authorization(authorization).Action(action).DescribeOneProjectRequest(describeOneProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeOneProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeOneProject`: DescribeOneProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeOneProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeOneProject&quot;]
 **describeOneProjectRequest** | [**DescribeOneProjectRequest**](DescribeOneProjectRequest.md) |  | 

### Return type

[**DescribeOneProject200Response**](DescribeOneProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePersonalExternalDepots

> DescribePersonalExternalDepots200Response DescribePersonalExternalDepots(ctx).Authorization(authorization).Action(action).DescribePersonalExternalDepotsRequest(describePersonalExternalDepotsRequest).Execute()

个人外部仓库获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePersonalExternalDepots" // string | Action (default to "DescribePersonalExternalDepots")
	describePersonalExternalDepotsRequest := *openapiclient.NewDescribePersonalExternalDepotsRequest("DepotType_example", int32(123)) // DescribePersonalExternalDepotsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePersonalExternalDepots(context.Background()).Authorization(authorization).Action(action).DescribePersonalExternalDepotsRequest(describePersonalExternalDepotsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePersonalExternalDepots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePersonalExternalDepots`: DescribePersonalExternalDepots200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePersonalExternalDepots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePersonalExternalDepotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePersonalExternalDepots&quot;]
 **describePersonalExternalDepotsRequest** | [**DescribePersonalExternalDepotsRequest**](DescribePersonalExternalDepotsRequest.md) |  | 

### Return type

[**DescribePersonalExternalDepots200Response**](DescribePersonalExternalDepots200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePinyin

> DescribePinyin200Response DescribePinyin(ctx).Authorization(authorization).Action(action).DescribePinyinRequest(describePinyinRequest).Execute()

汉字转拼音



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePinyin" // string | Action (default to "DescribePinyin")
	describePinyinRequest := *openapiclient.NewDescribePinyinRequest(false, "Style_example", "Value_example") // DescribePinyinRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePinyin(context.Background()).Authorization(authorization).Action(action).DescribePinyinRequest(describePinyinRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePinyin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePinyin`: DescribePinyin200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePinyin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePinyinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePinyin&quot;]
 **describePinyinRequest** | [**DescribePinyinRequest**](DescribePinyinRequest.md) |  | 

### Return type

[**DescribePinyin200Response**](DescribePinyin200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePoliciesOnResourceType

> DescribePoliciesOnResourceType200Response DescribePoliciesOnResourceType(ctx).Authorization(authorization).Action(action).DescribePoliciesOnResourceTypeRequest(describePoliciesOnResourceTypeRequest).Execute()

权限组列表查询（指定资源类型）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePoliciesOnResourceType" // string | Action (default to "DescribePoliciesOnResourceType")
	describePoliciesOnResourceTypeRequest := *openapiclient.NewDescribePoliciesOnResourceTypeRequest(*openapiclient.NewDescribePoliciesOnResourceTypeRequestFilter("ResourceType_example"), int32(123), int32(123)) // DescribePoliciesOnResourceTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePoliciesOnResourceType(context.Background()).Authorization(authorization).Action(action).DescribePoliciesOnResourceTypeRequest(describePoliciesOnResourceTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePoliciesOnResourceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePoliciesOnResourceType`: DescribePoliciesOnResourceType200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePoliciesOnResourceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePoliciesOnResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePoliciesOnResourceType&quot;]
 **describePoliciesOnResourceTypeRequest** | [**DescribePoliciesOnResourceTypeRequest**](DescribePoliciesOnResourceTypeRequest.md) |  | 

### Return type

[**DescribePoliciesOnResourceType200Response**](DescribePoliciesOnResourceType200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePolicy

> DescribePolicy200Response DescribePolicy(ctx).Authorization(authorization).Action(action).DescribePolicyRequest(describePolicyRequest).Execute()

权限组详情获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePolicy" // string | Action (default to "DescribePolicy")
	describePolicyRequest := *openapiclient.NewDescribePolicyRequest(false, int64(123)) // DescribePolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePolicy(context.Background()).Authorization(authorization).Action(action).DescribePolicyRequest(describePolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePolicy`: DescribePolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePolicy&quot;]
 **describePolicyRequest** | [**DescribePolicyRequest**](DescribePolicyRequest.md) |  | 

### Return type

[**DescribePolicy200Response**](DescribePolicy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePreSignUploadUrl

> DescribePreSignUploadUrl200Response DescribePreSignUploadUrl(ctx).Authorization(authorization).Action(action).DescribePreSignUploadUrlRequest(describePreSignUploadUrlRequest).Execute()

预签名URL获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePreSignUploadUrl" // string | Action (default to "DescribePreSignUploadUrl")
	describePreSignUploadUrlRequest := *openapiclient.NewDescribePreSignUploadUrlRequest("ContentType_example", "FileName_example", int64(123), int64(123), "ProjectName_example") // DescribePreSignUploadUrlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePreSignUploadUrl(context.Background()).Authorization(authorization).Action(action).DescribePreSignUploadUrlRequest(describePreSignUploadUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePreSignUploadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePreSignUploadUrl`: DescribePreSignUploadUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePreSignUploadUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePreSignUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePreSignUploadUrl&quot;]
 **describePreSignUploadUrlRequest** | [**DescribePreSignUploadUrlRequest**](DescribePreSignUploadUrlRequest.md) |  | 

### Return type

[**DescribePreSignUploadUrl200Response**](DescribePreSignUploadUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePredicatePolicyOnResource

> DescribePredicatePolicyOnResource200Response DescribePredicatePolicyOnResource(ctx).Authorization(authorization).Action(action).DescribePredicatePolicyOnResourceRequest(describePredicatePolicyOnResourceRequest).Execute()

资源权限判定模式获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePredicatePolicyOnResource" // string | Action (default to "DescribePredicatePolicyOnResource")
	describePredicatePolicyOnResourceRequest := *openapiclient.NewDescribePredicatePolicyOnResourceRequest(*openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example")) // DescribePredicatePolicyOnResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePredicatePolicyOnResource(context.Background()).Authorization(authorization).Action(action).DescribePredicatePolicyOnResourceRequest(describePredicatePolicyOnResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePredicatePolicyOnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePredicatePolicyOnResource`: DescribePredicatePolicyOnResource200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePredicatePolicyOnResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePredicatePolicyOnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePredicatePolicyOnResource&quot;]
 **describePredicatePolicyOnResourceRequest** | [**DescribePredicatePolicyOnResourceRequest**](DescribePredicatePolicyOnResourceRequest.md) |  | 

### Return type

[**DescribePredicatePolicyOnResource200Response**](DescribePredicatePolicyOnResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProgramProjects

> DescribeProgramProjects200Response DescribeProgramProjects(ctx).Authorization(authorization).Action(action).DescribeProgramProjectsRequest(describeProgramProjectsRequest).Execute()

项目集下项目列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProgramProjects" // string | Action (default to "DescribeProgramProjects")
	describeProgramProjectsRequest := *openapiclient.NewDescribeProgramProjectsRequest(int32(123)) // DescribeProgramProjectsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProgramProjects(context.Background()).Authorization(authorization).Action(action).DescribeProgramProjectsRequest(describeProgramProjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProgramProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProgramProjects`: DescribeProgramProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProgramProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProgramProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProgramProjects&quot;]
 **describeProgramProjectsRequest** | [**DescribeProgramProjectsRequest**](DescribeProgramProjectsRequest.md) |  | 

### Return type

[**DescribeProgramProjects200Response**](DescribeProgramProjects200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePrograms

> DescribePrograms200Response DescribePrograms(ctx).Authorization(authorization).Action(action).DescribeProgramsRequest(describeProgramsRequest).Execute()

项目集列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribePrograms" // string | Action (default to "DescribePrograms")
	describeProgramsRequest := *openapiclient.NewDescribeProgramsRequest(int32(1), int32(10)) // DescribeProgramsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribePrograms(context.Background()).Authorization(authorization).Action(action).DescribeProgramsRequest(describeProgramsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribePrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribePrograms`: DescribePrograms200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribePrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribePrograms&quot;]
 **describeProgramsRequest** | [**DescribeProgramsRequest**](DescribeProgramsRequest.md) |  | 

### Return type

[**DescribePrograms200Response**](DescribePrograms200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectAnnouncement

> CreateProjectAnnouncement200Response DescribeProjectAnnouncement(ctx).Authorization(authorization).Action(action).DescribeProjectAnnouncementRequest(describeProjectAnnouncementRequest).Execute()

项目公告查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectAnnouncement" // string | Action (default to "DescribeProjectAnnouncement")
	describeProjectAnnouncementRequest := *openapiclient.NewDescribeProjectAnnouncementRequest("ProjectName_example", int64(123)) // DescribeProjectAnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectAnnouncement(context.Background()).Authorization(authorization).Action(action).DescribeProjectAnnouncementRequest(describeProjectAnnouncementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectAnnouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectAnnouncement`: CreateProjectAnnouncement200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectAnnouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectAnnouncement&quot;]
 **describeProjectAnnouncementRequest** | [**DescribeProjectAnnouncementRequest**](DescribeProjectAnnouncementRequest.md) |  | 

### Return type

[**CreateProjectAnnouncement200Response**](CreateProjectAnnouncement200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectAnnouncements

> DescribeProjectAnnouncements200Response DescribeProjectAnnouncements(ctx).Authorization(authorization).Action(action).DescribeProjectAnnouncementsRequest(describeProjectAnnouncementsRequest).Execute()

项目公告列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectAnnouncements" // string | Action (default to "DescribeProjectAnnouncements")
	describeProjectAnnouncementsRequest := *openapiclient.NewDescribeProjectAnnouncementsRequest(int64(123), int64(123), "ProjectName_example") // DescribeProjectAnnouncementsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectAnnouncements(context.Background()).Authorization(authorization).Action(action).DescribeProjectAnnouncementsRequest(describeProjectAnnouncementsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectAnnouncements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectAnnouncements`: DescribeProjectAnnouncements200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectAnnouncements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectAnnouncementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectAnnouncements&quot;]
 **describeProjectAnnouncementsRequest** | [**DescribeProjectAnnouncementsRequest**](DescribeProjectAnnouncementsRequest.md) |  | 

### Return type

[**DescribeProjectAnnouncements200Response**](DescribeProjectAnnouncements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectByName

> DescribeOneProject200Response DescribeProjectByName(ctx).Authorization(authorization).Action(action).DescribeProjectByNameRequest(describeProjectByNameRequest).Execute()

项目查询(通过项目名称)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectByName" // string | Action (default to "DescribeProjectByName")
	describeProjectByNameRequest := *openapiclient.NewDescribeProjectByNameRequest("ProjectName_example") // DescribeProjectByNameRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectByName(context.Background()).Authorization(authorization).Action(action).DescribeProjectByNameRequest(describeProjectByNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectByName`: DescribeOneProject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectByName&quot;]
 **describeProjectByNameRequest** | [**DescribeProjectByNameRequest**](DescribeProjectByNameRequest.md) |  | 

### Return type

[**DescribeOneProject200Response**](DescribeOneProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectCredentials

> DescribeProjectCredentials200Response DescribeProjectCredentials(ctx).Authorization(authorization).Action(action).DescribeProjectCredentialsRequest(describeProjectCredentialsRequest).Execute()

项目凭据列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectCredentials" // string | Action (default to "DescribeProjectCredentials")
	describeProjectCredentialsRequest := *openapiclient.NewDescribeProjectCredentialsRequest(int32(123)) // DescribeProjectCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectCredentials(context.Background()).Authorization(authorization).Action(action).DescribeProjectCredentialsRequest(describeProjectCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectCredentials`: DescribeProjectCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectCredentials&quot;]
 **describeProjectCredentialsRequest** | [**DescribeProjectCredentialsRequest**](DescribeProjectCredentialsRequest.md) |  | 

### Return type

[**DescribeProjectCredentials200Response**](DescribeProjectCredentials200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectDepotBranches

> DescribeProjectDepotBranches200Response DescribeProjectDepotBranches(ctx).Authorization(authorization).Action(action).DescribeProjectDepotBranchesRequest(describeProjectDepotBranchesRequest).Execute()

仓库分支列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectDepotBranches" // string | Action (default to "DescribeProjectDepotBranches")
	describeProjectDepotBranchesRequest := *openapiclient.NewDescribeProjectDepotBranchesRequest("DepotType_example", int32(123), int32(123)) // DescribeProjectDepotBranchesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectDepotBranches(context.Background()).Authorization(authorization).Action(action).DescribeProjectDepotBranchesRequest(describeProjectDepotBranchesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectDepotBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectDepotBranches`: DescribeProjectDepotBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectDepotBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectDepotBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectDepotBranches&quot;]
 **describeProjectDepotBranchesRequest** | [**DescribeProjectDepotBranchesRequest**](DescribeProjectDepotBranchesRequest.md) |  | 

### Return type

[**DescribeProjectDepotBranches200Response**](DescribeProjectDepotBranches200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectDepotCommits

> DescribeProjectDepotBranches200Response DescribeProjectDepotCommits(ctx).Authorization(authorization).Action(action).DescribeProjectDepotCommitsRequest(describeProjectDepotCommitsRequest).Execute()

分支下的提交列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectDepotCommits" // string | Action (default to "DescribeProjectDepotCommits")
	describeProjectDepotCommitsRequest := *openapiclient.NewDescribeProjectDepotCommitsRequest("Branch_example", "DepotType_example", int32(123), int32(123)) // DescribeProjectDepotCommitsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectDepotCommits(context.Background()).Authorization(authorization).Action(action).DescribeProjectDepotCommitsRequest(describeProjectDepotCommitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectDepotCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectDepotCommits`: DescribeProjectDepotBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectDepotCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectDepotCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectDepotCommits&quot;]
 **describeProjectDepotCommitsRequest** | [**DescribeProjectDepotCommitsRequest**](DescribeProjectDepotCommitsRequest.md) |  | 

### Return type

[**DescribeProjectDepotBranches200Response**](DescribeProjectDepotBranches200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectDepotInfoList

> DescribeProjectDepotInfoList200Response DescribeProjectDepotInfoList(ctx).Authorization(authorization).Action(action).DescribeProjectDepotInfoListRequest(describeProjectDepotInfoListRequest).Execute()

仓库信息-查询项目下所有的仓库信息列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectDepotInfoList" // string | Action (default to "DescribeProjectDepotInfoList")
	describeProjectDepotInfoListRequest := *openapiclient.NewDescribeProjectDepotInfoListRequest(int64(123)) // DescribeProjectDepotInfoListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectDepotInfoList(context.Background()).Authorization(authorization).Action(action).DescribeProjectDepotInfoListRequest(describeProjectDepotInfoListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectDepotInfoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectDepotInfoList`: DescribeProjectDepotInfoList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectDepotInfoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectDepotInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectDepotInfoList&quot;]
 **describeProjectDepotInfoListRequest** | [**DescribeProjectDepotInfoListRequest**](DescribeProjectDepotInfoListRequest.md) |  | 

### Return type

[**DescribeProjectDepotInfoList200Response**](DescribeProjectDepotInfoList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectDepotTags

> DescribeProjectDepotBranches200Response DescribeProjectDepotTags(ctx).Authorization(authorization).Action(action).DescribeProjectDepotBranchesRequest(describeProjectDepotBranchesRequest).Execute()

仓库的标签列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectDepotTags" // string | Action (default to "DescribeProjectDepotTags")
	describeProjectDepotBranchesRequest := *openapiclient.NewDescribeProjectDepotBranchesRequest("DepotType_example", int32(123), int32(123)) // DescribeProjectDepotBranchesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectDepotTags(context.Background()).Authorization(authorization).Action(action).DescribeProjectDepotBranchesRequest(describeProjectDepotBranchesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectDepotTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectDepotTags`: DescribeProjectDepotBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectDepotTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectDepotTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectDepotTags&quot;]
 **describeProjectDepotBranchesRequest** | [**DescribeProjectDepotBranchesRequest**](DescribeProjectDepotBranchesRequest.md) |  | 

### Return type

[**DescribeProjectDepotBranches200Response**](DescribeProjectDepotBranches200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectDepots

> DescribeProjectDepots200Response DescribeProjectDepots(ctx).Authorization(authorization).Action(action).DescribePersonalExternalDepotsRequest(describePersonalExternalDepotsRequest).Execute()

项目仓库列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectDepots" // string | Action (default to "DescribeProjectDepots")
	describePersonalExternalDepotsRequest := *openapiclient.NewDescribePersonalExternalDepotsRequest("DepotType_example", int32(123)) // DescribePersonalExternalDepotsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectDepots(context.Background()).Authorization(authorization).Action(action).DescribePersonalExternalDepotsRequest(describePersonalExternalDepotsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectDepots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectDepots`: DescribeProjectDepots200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectDepots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectDepotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectDepots&quot;]
 **describePersonalExternalDepotsRequest** | [**DescribePersonalExternalDepotsRequest**](DescribePersonalExternalDepotsRequest.md) |  | 

### Return type

[**DescribeProjectDepots200Response**](DescribeProjectDepots200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectIssueFieldList

> DescribeProjectIssueFieldList200Response DescribeProjectIssueFieldList(ctx).Authorization(authorization).Action(action).DescribeProjectIssueFieldListRequest(describeProjectIssueFieldListRequest).Execute()

具体事项类型的属性列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectIssueFieldList" // string | Action (default to "DescribeProjectIssueFieldList")
	describeProjectIssueFieldListRequest := *openapiclient.NewDescribeProjectIssueFieldListRequest("IssueType_example", "ProjectName_example") // DescribeProjectIssueFieldListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectIssueFieldList(context.Background()).Authorization(authorization).Action(action).DescribeProjectIssueFieldListRequest(describeProjectIssueFieldListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectIssueFieldList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectIssueFieldList`: DescribeProjectIssueFieldList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectIssueFieldList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectIssueFieldListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectIssueFieldList&quot;]
 **describeProjectIssueFieldListRequest** | [**DescribeProjectIssueFieldListRequest**](DescribeProjectIssueFieldListRequest.md) |  | 

### Return type

[**DescribeProjectIssueFieldList200Response**](DescribeProjectIssueFieldList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectIssueStatusList

> DescribeProjectIssueStatusList200Response DescribeProjectIssueStatusList(ctx).Authorization(authorization).Action(action).DescribeProjectIssueFieldListRequest(describeProjectIssueFieldListRequest).Execute()

具体事项类型的状态列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectIssueStatusList" // string | Action (default to "DescribeProjectIssueStatusList")
	describeProjectIssueFieldListRequest := *openapiclient.NewDescribeProjectIssueFieldListRequest("IssueType_example", "ProjectName_example") // DescribeProjectIssueFieldListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectIssueStatusList(context.Background()).Authorization(authorization).Action(action).DescribeProjectIssueFieldListRequest(describeProjectIssueFieldListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectIssueStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectIssueStatusList`: DescribeProjectIssueStatusList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectIssueStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectIssueStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectIssueStatusList&quot;]
 **describeProjectIssueFieldListRequest** | [**DescribeProjectIssueFieldListRequest**](DescribeProjectIssueFieldListRequest.md) |  | 

### Return type

[**DescribeProjectIssueStatusList200Response**](DescribeProjectIssueStatusList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectIssueTypeList

> DescribeProjectIssueTypeList200Response DescribeProjectIssueTypeList(ctx).Authorization(authorization).Action(action).DescribeProjectIssueTypeListRequest(describeProjectIssueTypeListRequest).Execute()

项目事项类型列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectIssueTypeList" // string | Action (default to "DescribeProjectIssueTypeList")
	describeProjectIssueTypeListRequest := *openapiclient.NewDescribeProjectIssueTypeListRequest("ProjectName_example") // DescribeProjectIssueTypeListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectIssueTypeList(context.Background()).Authorization(authorization).Action(action).DescribeProjectIssueTypeListRequest(describeProjectIssueTypeListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectIssueTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectIssueTypeList`: DescribeProjectIssueTypeList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectIssueTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectIssueTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectIssueTypeList&quot;]
 **describeProjectIssueTypeListRequest** | [**DescribeProjectIssueTypeListRequest**](DescribeProjectIssueTypeListRequest.md) |  | 

### Return type

[**DescribeProjectIssueTypeList200Response**](DescribeProjectIssueTypeList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectLabels

> DescribeProjectLabels200Response DescribeProjectLabels(ctx).Authorization(authorization).Action(action).DescribeProjectLabelsRequest(describeProjectLabelsRequest).Execute()

项目列表查询-指定项目标签



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectLabels" // string | Action (default to "DescribeProjectLabels")
	describeProjectLabelsRequest := *openapiclient.NewDescribeProjectLabelsRequest() // DescribeProjectLabelsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectLabels(context.Background()).Authorization(authorization).Action(action).DescribeProjectLabelsRequest(describeProjectLabelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectLabels`: DescribeProjectLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectLabels&quot;]
 **describeProjectLabelsRequest** | [**DescribeProjectLabelsRequest**](DescribeProjectLabelsRequest.md) |  | 

### Return type

[**DescribeProjectLabels200Response**](DescribeProjectLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectMemberPrincipals

> DescribeProjectMemberPrincipals200Response DescribeProjectMemberPrincipals(ctx).Authorization(authorization).Action(action).DescribeProjectMemberPrincipalsRequest(describeProjectMemberPrincipalsRequest).Execute()

项目成员主体查询(包含用户组、部门、成员)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectMemberPrincipals" // string | Action (default to "DescribeProjectMemberPrincipals")
	describeProjectMemberPrincipalsRequest := *openapiclient.NewDescribeProjectMemberPrincipalsRequest(int32(123), int32(123), int32(123)) // DescribeProjectMemberPrincipalsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectMemberPrincipals(context.Background()).Authorization(authorization).Action(action).DescribeProjectMemberPrincipalsRequest(describeProjectMemberPrincipalsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectMemberPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectMemberPrincipals`: DescribeProjectMemberPrincipals200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectMemberPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectMemberPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectMemberPrincipals&quot;]
 **describeProjectMemberPrincipalsRequest** | [**DescribeProjectMemberPrincipalsRequest**](DescribeProjectMemberPrincipalsRequest.md) |  | 

### Return type

[**DescribeProjectMemberPrincipals200Response**](DescribeProjectMemberPrincipals200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectMembers

> DescribeProjectMembers200Response DescribeProjectMembers(ctx).Authorization(authorization).Action(action).DescribeProjectMembersRequest(describeProjectMembersRequest).Execute()

项目成员列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectMembers" // string | Action (default to "DescribeProjectMembers")
	describeProjectMembersRequest := *openapiclient.NewDescribeProjectMembersRequest(int32(123), int32(123), int32(123)) // DescribeProjectMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectMembers(context.Background()).Authorization(authorization).Action(action).DescribeProjectMembersRequest(describeProjectMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectMembers`: DescribeProjectMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectMembers&quot;]
 **describeProjectMembersRequest** | [**DescribeProjectMembersRequest**](DescribeProjectMembersRequest.md) |  | 

### Return type

[**DescribeProjectMembers200Response**](DescribeProjectMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectMergeRequests

> DescribeSelfMergeRequests200Response DescribeProjectMergeRequests(ctx).Authorization(authorization).Action(action).DescribeProjectMergeRequestsRequest(describeProjectMergeRequestsRequest).Execute()

合并请求-获取项目下的合并请求列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectMergeRequests" // string | Action (default to "DescribeProjectMergeRequests")
	describeProjectMergeRequestsRequest := *openapiclient.NewDescribeProjectMergeRequestsRequest(int64(123)) // DescribeProjectMergeRequestsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectMergeRequests(context.Background()).Authorization(authorization).Action(action).DescribeProjectMergeRequestsRequest(describeProjectMergeRequestsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectMergeRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectMergeRequests`: DescribeSelfMergeRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectMergeRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectMergeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectMergeRequests&quot;]
 **describeProjectMergeRequestsRequest** | [**DescribeProjectMergeRequestsRequest**](DescribeProjectMergeRequestsRequest.md) |  | 

### Return type

[**DescribeSelfMergeRequests200Response**](DescribeSelfMergeRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectRoles

> DescribeProjectRoles200Response DescribeProjectRoles(ctx).Authorization(authorization).Action(action).DescribeProjectRolesRequest(describeProjectRolesRequest).Execute()

项目用户组查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectRoles" // string | Action (default to "DescribeProjectRoles")
	describeProjectRolesRequest := *openapiclient.NewDescribeProjectRolesRequest(int32(123)) // DescribeProjectRolesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectRoles(context.Background()).Authorization(authorization).Action(action).DescribeProjectRolesRequest(describeProjectRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectRoles`: DescribeProjectRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectRoles&quot;]
 **describeProjectRolesRequest** | [**DescribeProjectRolesRequest**](DescribeProjectRolesRequest.md) |  | 

### Return type

[**DescribeProjectRoles200Response**](DescribeProjectRoles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProjectsByFeature

> DescribeProjectsByFeature200Response DescribeProjectsByFeature(ctx).Authorization(authorization).Action(action).DescribeProjectsByFeatureRequest(describeProjectsByFeatureRequest).Execute()

项目查询（通过一级菜单名）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProjectsByFeature" // string | Action (default to "DescribeProjectsByFeature")
	describeProjectsByFeatureRequest := *openapiclient.NewDescribeProjectsByFeatureRequest("MenuName_example") // DescribeProjectsByFeatureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProjectsByFeature(context.Background()).Authorization(authorization).Action(action).DescribeProjectsByFeatureRequest(describeProjectsByFeatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProjectsByFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProjectsByFeature`: DescribeProjectsByFeature200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProjectsByFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProjectsByFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProjectsByFeature&quot;]
 **describeProjectsByFeatureRequest** | [**DescribeProjectsByFeatureRequest**](DescribeProjectsByFeatureRequest.md) |  | 

### Return type

[**DescribeProjectsByFeature200Response**](DescribeProjectsByFeature200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProtectedBranch

> DescribeProtectedBranch200Response DescribeProtectedBranch(ctx).Authorization(authorization).Action(action).DescribeProtectedBranchRequest(describeProtectedBranchRequest).Execute()

仓库设置-查询保护分支详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProtectedBranch" // string | Action (default to "DescribeProtectedBranch")
	describeProtectedBranchRequest := *openapiclient.NewDescribeProtectedBranchRequest("BranchName_example", int64(123)) // DescribeProtectedBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProtectedBranch(context.Background()).Authorization(authorization).Action(action).DescribeProtectedBranchRequest(describeProtectedBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProtectedBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProtectedBranch`: DescribeProtectedBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProtectedBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProtectedBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProtectedBranch&quot;]
 **describeProtectedBranchRequest** | [**DescribeProtectedBranchRequest**](DescribeProtectedBranchRequest.md) |  | 

### Return type

[**DescribeProtectedBranch200Response**](DescribeProtectedBranch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProtectedBranchMembers

> DescribeProtectedBranchMembers200Response DescribeProtectedBranchMembers(ctx).Authorization(authorization).Action(action).DescribeProtectedBranchMembersRequest(describeProtectedBranchMembersRequest).Execute()

仓库设置-查询保护分支成员



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProtectedBranchMembers" // string | Action (default to "DescribeProtectedBranchMembers")
	describeProtectedBranchMembersRequest := *openapiclient.NewDescribeProtectedBranchMembersRequest("BranchName_example", int64(123)) // DescribeProtectedBranchMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProtectedBranchMembers(context.Background()).Authorization(authorization).Action(action).DescribeProtectedBranchMembersRequest(describeProtectedBranchMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProtectedBranchMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProtectedBranchMembers`: DescribeProtectedBranchMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProtectedBranchMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProtectedBranchMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProtectedBranchMembers&quot;]
 **describeProtectedBranchMembersRequest** | [**DescribeProtectedBranchMembersRequest**](DescribeProtectedBranchMembersRequest.md) |  | 

### Return type

[**DescribeProtectedBranchMembers200Response**](DescribeProtectedBranchMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeProtectedBranches

> DescribeProtectedBranches200Response DescribeProtectedBranches(ctx).Authorization(authorization).Action(action).DescribeProtectedBranchesRequest(describeProtectedBranchesRequest).Execute()

仓库设置-查询保护分支列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeProtectedBranches" // string | Action (default to "DescribeProtectedBranches")
	describeProtectedBranchesRequest := *openapiclient.NewDescribeProtectedBranchesRequest(int64(123), "DepotPath_example") // DescribeProtectedBranchesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeProtectedBranches(context.Background()).Authorization(authorization).Action(action).DescribeProtectedBranchesRequest(describeProtectedBranchesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeProtectedBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeProtectedBranches`: DescribeProtectedBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeProtectedBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeProtectedBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeProtectedBranches&quot;]
 **describeProtectedBranchesRequest** | [**DescribeProtectedBranchesRequest**](DescribeProtectedBranchesRequest.md) |  | 

### Return type

[**DescribeProtectedBranches200Response**](DescribeProtectedBranches200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeRelatedCaseList

> DescribeRelatedCaseList200Response DescribeRelatedCaseList(ctx).Authorization(authorization).Action(action).DescribeRelatedCaseListRequest(describeRelatedCaseListRequest).Execute()

事项关联的测试用例查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeRelatedCaseList" // string | Action (default to "DescribeRelatedCaseList")
	describeRelatedCaseListRequest := *openapiclient.NewDescribeRelatedCaseListRequest(int64(123), "ProjectName_example") // DescribeRelatedCaseListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeRelatedCaseList(context.Background()).Authorization(authorization).Action(action).DescribeRelatedCaseListRequest(describeRelatedCaseListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeRelatedCaseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeRelatedCaseList`: DescribeRelatedCaseList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeRelatedCaseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRelatedCaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeRelatedCaseList&quot;]
 **describeRelatedCaseListRequest** | [**DescribeRelatedCaseListRequest**](DescribeRelatedCaseListRequest.md) |  | 

### Return type

[**DescribeRelatedCaseList200Response**](DescribeRelatedCaseList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeRelease

> ModifyRelease200Response DescribeRelease(ctx).Authorization(authorization).Action(action).DescribeReleaseRequest(describeReleaseRequest).Execute()

版本详情查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeRelease" // string | Action (default to "DescribeRelease")
	describeReleaseRequest := *openapiclient.NewDescribeReleaseRequest("ProjectName_example", int64(123)) // DescribeReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeRelease(context.Background()).Authorization(authorization).Action(action).DescribeReleaseRequest(describeReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeRelease`: ModifyRelease200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeRelease&quot;]
 **describeReleaseRequest** | [**DescribeReleaseRequest**](DescribeReleaseRequest.md) |  | 

### Return type

[**ModifyRelease200Response**](ModifyRelease200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeReleaseIssueList

> DescribeReleaseIssueList200Response DescribeReleaseIssueList(ctx).Authorization(authorization).Action(action).DescribeReleaseIssueListRequest(describeReleaseIssueListRequest).Execute()

版本发布范围查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeReleaseIssueList" // string | Action (default to "DescribeReleaseIssueList")
	describeReleaseIssueListRequest := *openapiclient.NewDescribeReleaseIssueListRequest("ProjectName_example", int64(123)) // DescribeReleaseIssueListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeReleaseIssueList(context.Background()).Authorization(authorization).Action(action).DescribeReleaseIssueListRequest(describeReleaseIssueListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeReleaseIssueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeReleaseIssueList`: DescribeReleaseIssueList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeReleaseIssueList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeReleaseIssueListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeReleaseIssueList&quot;]
 **describeReleaseIssueListRequest** | [**DescribeReleaseIssueListRequest**](DescribeReleaseIssueListRequest.md) |  | 

### Return type

[**DescribeReleaseIssueList200Response**](DescribeReleaseIssueList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeReleaseList

> DescribeReleaseList200Response DescribeReleaseList(ctx).Authorization(authorization).Action(action).DescribeReleaseListRequest(describeReleaseListRequest).Execute()

版本列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeReleaseList" // string | Action (default to "DescribeReleaseList")
	describeReleaseListRequest := *openapiclient.NewDescribeReleaseListRequest("ProjectName_example") // DescribeReleaseListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeReleaseList(context.Background()).Authorization(authorization).Action(action).DescribeReleaseListRequest(describeReleaseListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeReleaseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeReleaseList`: DescribeReleaseList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeReleaseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeReleaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeReleaseList&quot;]
 **describeReleaseListRequest** | [**DescribeReleaseListRequest**](DescribeReleaseListRequest.md) |  | 

### Return type

[**DescribeReleaseList200Response**](DescribeReleaseList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeReport

> DescribeReport200Response DescribeReport(ctx).Authorization(authorization).Action(action).DeleteReportRequest(deleteReportRequest).Execute()

测试报告详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeReport" // string | Action (default to "DescribeReport")
	deleteReportRequest := *openapiclient.NewDeleteReportRequest("ProjectName_example", int32(123)) // DeleteReportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeReport(context.Background()).Authorization(authorization).Action(action).DeleteReportRequest(deleteReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeReport`: DescribeReport200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeReport&quot;]
 **deleteReportRequest** | [**DeleteReportRequest**](DeleteReportRequest.md) |  | 

### Return type

[**DescribeReport200Response**](DescribeReport200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeReportList

> DescribeReportList200Response DescribeReportList(ctx).Authorization(authorization).Action(action).DescribeReportListRequest(describeReportListRequest).Execute()

测试报告列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeReportList" // string | Action (default to "DescribeReportList")
	describeReportListRequest := *openapiclient.NewDescribeReportListRequest("ProjectName_example") // DescribeReportListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeReportList(context.Background()).Authorization(authorization).Action(action).DescribeReportListRequest(describeReportListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeReportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeReportList`: DescribeReportList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeReportList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeReportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeReportList&quot;]
 **describeReportListRequest** | [**DescribeReportListRequest**](DescribeReportListRequest.md) |  | 

### Return type

[**DescribeReportList200Response**](DescribeReportList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeRequirementDefectRelation

> DescribeRequirementDefectRelation200Response DescribeRequirementDefectRelation(ctx).Authorization(authorization).Action(action).DescribeRequirementDefectRelationRequest(describeRequirementDefectRelationRequest).Execute()

需求关联缺陷列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeRequirementDefectRelation" // string | Action (default to "DescribeRequirementDefectRelation")
	describeRequirementDefectRelationRequest := *openapiclient.NewDescribeRequirementDefectRelationRequest("ProjectName_example", int64(123)) // DescribeRequirementDefectRelationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeRequirementDefectRelation(context.Background()).Authorization(authorization).Action(action).DescribeRequirementDefectRelationRequest(describeRequirementDefectRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeRequirementDefectRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeRequirementDefectRelation`: DescribeRequirementDefectRelation200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeRequirementDefectRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRequirementDefectRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeRequirementDefectRelation&quot;]
 **describeRequirementDefectRelationRequest** | [**DescribeRequirementDefectRelationRequest**](DescribeRequirementDefectRelationRequest.md) |  | 

### Return type

[**DescribeRequirementDefectRelation200Response**](DescribeRequirementDefectRelation200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeRequirementTestCaseList

> DescribeTestCaseList200Response DescribeRequirementTestCaseList(ctx).Authorization(authorization).Action(action).DescribeRequirementTestCaseListRequest(describeRequirementTestCaseListRequest).Execute()

需求关联的测试用例列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeRequirementTestCaseList" // string | Action (default to "DescribeRequirementTestCaseList")
	describeRequirementTestCaseListRequest := *openapiclient.NewDescribeRequirementTestCaseListRequest(int32(123), "ProjectName_example") // DescribeRequirementTestCaseListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeRequirementTestCaseList(context.Background()).Authorization(authorization).Action(action).DescribeRequirementTestCaseListRequest(describeRequirementTestCaseListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeRequirementTestCaseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeRequirementTestCaseList`: DescribeTestCaseList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeRequirementTestCaseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRequirementTestCaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeRequirementTestCaseList&quot;]
 **describeRequirementTestCaseListRequest** | [**DescribeRequirementTestCaseListRequest**](DescribeRequirementTestCaseListRequest.md) |  | 

### Return type

[**DescribeTestCaseList200Response**](DescribeTestCaseList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeResourceReferences

> DescribeResourceReferences200Response DescribeResourceReferences(ctx).Authorization(authorization).Action(action).DescribeResourceReferencesRequest(describeResourceReferencesRequest).Execute()

资源引用的资源列表，如 开发任务中引用了多个需求，获取任务引用的需求列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeResourceReferences" // string | Action (default to "DescribeResourceReferences")
	describeResourceReferencesRequest := *openapiclient.NewDescribeResourceReferencesRequest("ProjectName_example", int64(123)) // DescribeResourceReferencesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeResourceReferences(context.Background()).Authorization(authorization).Action(action).DescribeResourceReferencesRequest(describeResourceReferencesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeResourceReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeResourceReferences`: DescribeResourceReferences200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeResourceReferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeResourceReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeResourceReferences&quot;]
 **describeResourceReferencesRequest** | [**DescribeResourceReferencesRequest**](DescribeResourceReferencesRequest.md) |  | 

### Return type

[**DescribeResourceReferences200Response**](DescribeResourceReferences200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeResourceReferencesCited

> DescribeResourceReferencesCited200Response DescribeResourceReferencesCited(ctx).Authorization(authorization).Action(action).DescribeResourceReferencesCitingRequest(describeResourceReferencesCitingRequest).Execute()

被引用资源列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeResourceReferencesCited" // string | Action (default to "DescribeResourceReferencesCited")
	describeResourceReferencesCitingRequest := *openapiclient.NewDescribeResourceReferencesCitingRequest(int64(123), "ResourceCode_example", int64(123)) // DescribeResourceReferencesCitingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeResourceReferencesCited(context.Background()).Authorization(authorization).Action(action).DescribeResourceReferencesCitingRequest(describeResourceReferencesCitingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeResourceReferencesCited``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeResourceReferencesCited`: DescribeResourceReferencesCited200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeResourceReferencesCited`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeResourceReferencesCitedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeResourceReferencesCited&quot;]
 **describeResourceReferencesCitingRequest** | [**DescribeResourceReferencesCitingRequest**](DescribeResourceReferencesCitingRequest.md) |  | 

### Return type

[**DescribeResourceReferencesCited200Response**](DescribeResourceReferencesCited200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeResourceReferencesCiting

> DescribeResourceReferencesCiting200Response DescribeResourceReferencesCiting(ctx).Authorization(authorization).Action(action).DescribeResourceReferencesCitingRequest(describeResourceReferencesCitingRequest).Execute()

引用资源列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeResourceReferencesCiting" // string | Action (default to "DescribeResourceReferencesCiting")
	describeResourceReferencesCitingRequest := *openapiclient.NewDescribeResourceReferencesCitingRequest(int64(123), "ResourceCode_example", int64(123)) // DescribeResourceReferencesCitingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeResourceReferencesCiting(context.Background()).Authorization(authorization).Action(action).DescribeResourceReferencesCitingRequest(describeResourceReferencesCitingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeResourceReferencesCiting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeResourceReferencesCiting`: DescribeResourceReferencesCiting200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeResourceReferencesCiting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeResourceReferencesCitingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeResourceReferencesCiting&quot;]
 **describeResourceReferencesCitingRequest** | [**DescribeResourceReferencesCitingRequest**](DescribeResourceReferencesCitingRequest.md) |  | 

### Return type

[**DescribeResourceReferencesCiting200Response**](DescribeResourceReferencesCiting200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeResourceScopeListOnPolicy

> DescribeResourceScopeListOnPolicy200Response DescribeResourceScopeListOnPolicy(ctx).Authorization(authorization).Action(action).DescribeResourceScopeListOnPolicyRequest(describeResourceScopeListOnPolicyRequest).Execute()

权限组可用资源范围分页查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeResourceScopeListOnPolicy" // string | Action (default to "DescribeResourceScopeListOnPolicy")
	describeResourceScopeListOnPolicyRequest := *openapiclient.NewDescribeResourceScopeListOnPolicyRequest(int32(123), int32(123), int64(123)) // DescribeResourceScopeListOnPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeResourceScopeListOnPolicy(context.Background()).Authorization(authorization).Action(action).DescribeResourceScopeListOnPolicyRequest(describeResourceScopeListOnPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeResourceScopeListOnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeResourceScopeListOnPolicy`: DescribeResourceScopeListOnPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeResourceScopeListOnPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeResourceScopeListOnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeResourceScopeListOnPolicy&quot;]
 **describeResourceScopeListOnPolicyRequest** | [**DescribeResourceScopeListOnPolicyRequest**](DescribeResourceScopeListOnPolicyRequest.md) |  | 

### Return type

[**DescribeResourceScopeListOnPolicy200Response**](DescribeResourceScopeListOnPolicy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeSelfMergeRequests

> DescribeSelfMergeRequests200Response DescribeSelfMergeRequests(ctx).Authorization(authorization).Action(action).DescribeSelfMergeRequestsRequest(describeSelfMergeRequestsRequest).Execute()

合并请求-获取自己的合并请求列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeSelfMergeRequests" // string | Action (default to "DescribeSelfMergeRequests")
	describeSelfMergeRequestsRequest := *openapiclient.NewDescribeSelfMergeRequestsRequest() // DescribeSelfMergeRequestsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeSelfMergeRequests(context.Background()).Authorization(authorization).Action(action).DescribeSelfMergeRequestsRequest(describeSelfMergeRequestsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeSelfMergeRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeSelfMergeRequests`: DescribeSelfMergeRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeSelfMergeRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSelfMergeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeSelfMergeRequests&quot;]
 **describeSelfMergeRequestsRequest** | [**DescribeSelfMergeRequestsRequest**](DescribeSelfMergeRequestsRequest.md) |  | 

### Return type

[**DescribeSelfMergeRequests200Response**](DescribeSelfMergeRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeSingeMergeRequestNotes

> DescribeSingeMergeRequestNotes200Response DescribeSingeMergeRequestNotes(ctx).Authorization(authorization).Action(action).DescribeSingeMergeRequestNotesRequest(describeSingeMergeRequestNotesRequest).Execute()

合并请求-获取单个合并请求行评论和改动文件diff行评论



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeSingeMergeRequestNotes" // string | Action (default to "DescribeSingeMergeRequestNotes")
	describeSingeMergeRequestNotesRequest := *openapiclient.NewDescribeSingeMergeRequestNotesRequest("DepotPath_example", int32(123)) // DescribeSingeMergeRequestNotesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeSingeMergeRequestNotes(context.Background()).Authorization(authorization).Action(action).DescribeSingeMergeRequestNotesRequest(describeSingeMergeRequestNotesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeSingeMergeRequestNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeSingeMergeRequestNotes`: DescribeSingeMergeRequestNotes200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeSingeMergeRequestNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSingeMergeRequestNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeSingeMergeRequestNotes&quot;]
 **describeSingeMergeRequestNotesRequest** | [**DescribeSingeMergeRequestNotesRequest**](DescribeSingeMergeRequestNotesRequest.md) |  | 

### Return type

[**DescribeSingeMergeRequestNotes200Response**](DescribeSingeMergeRequestNotes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeSshKey

> DescribeSshKey200Response DescribeSshKey(ctx).Authorization(authorization).Action(action).Body(body).Execute()

仓库设置-获取当前用户所有SSH公钥



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeSshKey" // string | Action (default to "DescribeSshKey")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeSshKey(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeSshKey`: DescribeSshKey200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeSshKey&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeSshKey200Response**](DescribeSshKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeSubIssueList

> DescribeSubIssueList200Response DescribeSubIssueList(ctx).Authorization(authorization).Action(action).DescribeSubIssueListRequest(describeSubIssueListRequest).Execute()

子事项列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeSubIssueList" // string | Action (default to "DescribeSubIssueList")
	describeSubIssueListRequest := *openapiclient.NewDescribeSubIssueListRequest(int64(123), "ProjectName_example") // DescribeSubIssueListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeSubIssueList(context.Background()).Authorization(authorization).Action(action).DescribeSubIssueListRequest(describeSubIssueListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeSubIssueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeSubIssueList`: DescribeSubIssueList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeSubIssueList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeSubIssueListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeSubIssueList&quot;]
 **describeSubIssueListRequest** | [**DescribeSubIssueListRequest**](DescribeSubIssueListRequest.md) |  | 

### Return type

[**DescribeSubIssueList200Response**](DescribeSubIssueList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeam

> DescribeTeam200Response DescribeTeam(ctx).Authorization(authorization).Action(action).Body(body).Execute()

团队信息查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeam" // string | Action (default to "DescribeTeam")
	body := interface{}(987) // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeam(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeam`: DescribeTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeam&quot;]
 **body** | **interface{}** |  | 

### Return type

[**DescribeTeam200Response**](DescribeTeam200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamAdminMembers

> DescribeTeamAdminMembers200Response DescribeTeamAdminMembers(ctx).Authorization(authorization).Action(action).DescribeTeamAdminMembersRequest(describeTeamAdminMembersRequest).Execute()

团队管理员查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamAdminMembers" // string | Action (default to "DescribeTeamAdminMembers")
	describeTeamAdminMembersRequest := *openapiclient.NewDescribeTeamAdminMembersRequest(int32(123), int32(123)) // DescribeTeamAdminMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamAdminMembers(context.Background()).Authorization(authorization).Action(action).DescribeTeamAdminMembersRequest(describeTeamAdminMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamAdminMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamAdminMembers`: DescribeTeamAdminMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamAdminMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamAdminMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamAdminMembers&quot;]
 **describeTeamAdminMembersRequest** | [**DescribeTeamAdminMembersRequest**](DescribeTeamAdminMembersRequest.md) |  | 

### Return type

[**DescribeTeamAdminMembers200Response**](DescribeTeamAdminMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamArtifacts

> DescribeTeamArtifacts200Response DescribeTeamArtifacts(ctx).Authorization(authorization).Action(action).DescribeTeamArtifactsRequest(describeTeamArtifactsRequest).Execute()

制品列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamArtifacts" // string | Action (default to "DescribeTeamArtifacts")
	describeTeamArtifactsRequest := *openapiclient.NewDescribeTeamArtifactsRequest() // DescribeTeamArtifactsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamArtifacts(context.Background()).Authorization(authorization).Action(action).DescribeTeamArtifactsRequest(describeTeamArtifactsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamArtifacts`: DescribeTeamArtifacts200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamArtifacts&quot;]
 **describeTeamArtifactsRequest** | [**DescribeTeamArtifactsRequest**](DescribeTeamArtifactsRequest.md) |  | 

### Return type

[**DescribeTeamArtifacts200Response**](DescribeTeamArtifacts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamDepotInfoList

> DescribeTeamDepotInfoList200Response DescribeTeamDepotInfoList(ctx).Authorization(authorization).Action(action).DescribeTeamDepotInfoListRequest(describeTeamDepotInfoListRequest).Execute()

仓库信息-获取团队下仓库列表，仅团队所有者或团队管理员可以调用。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamDepotInfoList" // string | Action (default to "DescribeTeamDepotInfoList")
	describeTeamDepotInfoListRequest := *openapiclient.NewDescribeTeamDepotInfoListRequest(int64(123), int64(123)) // DescribeTeamDepotInfoListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamDepotInfoList(context.Background()).Authorization(authorization).Action(action).DescribeTeamDepotInfoListRequest(describeTeamDepotInfoListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamDepotInfoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamDepotInfoList`: DescribeTeamDepotInfoList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamDepotInfoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamDepotInfoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamDepotInfoList&quot;]
 **describeTeamDepotInfoListRequest** | [**DescribeTeamDepotInfoListRequest**](DescribeTeamDepotInfoListRequest.md) |  | 

### Return type

[**DescribeTeamDepotInfoList200Response**](DescribeTeamDepotInfoList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamDepots

> DescribeTeamDepots200Response DescribeTeamDepots(ctx).Authorization(authorization).Action(action).DescribeTeamDepotsRequest(describeTeamDepotsRequest).Execute()

团队下可访问的所有仓库列表获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamDepots" // string | Action (default to "DescribeTeamDepots")
	describeTeamDepotsRequest := *openapiclient.NewDescribeTeamDepotsRequest("DepotType_example") // DescribeTeamDepotsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamDepots(context.Background()).Authorization(authorization).Action(action).DescribeTeamDepotsRequest(describeTeamDepotsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamDepots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamDepots`: DescribeTeamDepots200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamDepots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamDepotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamDepots&quot;]
 **describeTeamDepotsRequest** | [**DescribeTeamDepotsRequest**](DescribeTeamDepotsRequest.md) |  | 

### Return type

[**DescribeTeamDepots200Response**](DescribeTeamDepots200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamIssueTypeList

> DescribeTeamIssueTypeList200Response DescribeTeamIssueTypeList(ctx).Authorization(authorization).Action(action).Body(body).Execute()

企业事项类型列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamIssueTypeList" // string | Action (default to "DescribeTeamIssueTypeList")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamIssueTypeList(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamIssueTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamIssueTypeList`: DescribeTeamIssueTypeList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamIssueTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamIssueTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamIssueTypeList&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeTeamIssueTypeList200Response**](DescribeTeamIssueTypeList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamMember

> DescribeTeamMember200Response DescribeTeamMember(ctx).Authorization(authorization).Action(action).DescribeTeamMemberRequest(describeTeamMemberRequest).Execute()

团队成员信息查询（通过用户 ID）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamMember" // string | Action (default to "DescribeTeamMember")
	describeTeamMemberRequest := *openapiclient.NewDescribeTeamMemberRequest(int32(123)) // DescribeTeamMemberRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamMember(context.Background()).Authorization(authorization).Action(action).DescribeTeamMemberRequest(describeTeamMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamMember`: DescribeTeamMember200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamMember&quot;]
 **describeTeamMemberRequest** | [**DescribeTeamMemberRequest**](DescribeTeamMemberRequest.md) |  | 

### Return type

[**DescribeTeamMember200Response**](DescribeTeamMember200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamMemberByEmail

> DescribeTeamMemberByEmail200Response DescribeTeamMemberByEmail(ctx).Authorization(authorization).Action(action).DescribeTeamMemberByEmailRequest(describeTeamMemberByEmailRequest).Execute()

团队成员信息查询（通过用户 Email）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamMemberByEmail" // string | Action (default to "DescribeTeamMemberByEmail")
	describeTeamMemberByEmailRequest := *openapiclient.NewDescribeTeamMemberByEmailRequest("Email_example") // DescribeTeamMemberByEmailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamMemberByEmail(context.Background()).Authorization(authorization).Action(action).DescribeTeamMemberByEmailRequest(describeTeamMemberByEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamMemberByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamMemberByEmail`: DescribeTeamMemberByEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamMemberByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamMemberByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamMemberByEmail&quot;]
 **describeTeamMemberByEmailRequest** | [**DescribeTeamMemberByEmailRequest**](DescribeTeamMemberByEmailRequest.md) |  | 

### Return type

[**DescribeTeamMemberByEmail200Response**](DescribeTeamMemberByEmail200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTeamMembers

> DescribeTeamMembers200Response DescribeTeamMembers(ctx).Authorization(authorization).Action(action).DescribeTeamMembersRequest(describeTeamMembersRequest).Execute()

团队成员列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTeamMembers" // string | Action (default to "DescribeTeamMembers")
	describeTeamMembersRequest := *openapiclient.NewDescribeTeamMembersRequest(int32(123), int32(123)) // DescribeTeamMembersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTeamMembers(context.Background()).Authorization(authorization).Action(action).DescribeTeamMembersRequest(describeTeamMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTeamMembers`: DescribeTeamMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTeamMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTeamMembers&quot;]
 **describeTeamMembersRequest** | [**DescribeTeamMembersRequest**](DescribeTeamMembersRequest.md) |  | 

### Return type

[**DescribeTeamMembers200Response**](DescribeTeamMembers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTest

> DescribeTest200Response DescribeTest(ctx).Authorization(authorization).Action(action).DescribeTestRequest(describeTestRequest).Execute()

测试任务详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTest" // string | Action (default to "DescribeTest")
	describeTestRequest := *openapiclient.NewDescribeTestRequest("ProjectName_example") // DescribeTestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTest(context.Background()).Authorization(authorization).Action(action).DescribeTestRequest(describeTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTest`: DescribeTest200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTest&quot;]
 **describeTestRequest** | [**DescribeTestRequest**](DescribeTestRequest.md) |  | 

### Return type

[**DescribeTest200Response**](DescribeTest200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestCase

> ModifyTestCase200Response DescribeTestCase(ctx).Authorization(authorization).Action(action).DeleteTestCaseRequest(deleteTestCaseRequest).Execute()

测试用例详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestCase" // string | Action (default to "DescribeTestCase")
	deleteTestCaseRequest := *openapiclient.NewDeleteTestCaseRequest(int32(123), "ProjectName_example") // DeleteTestCaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestCase(context.Background()).Authorization(authorization).Action(action).DeleteTestCaseRequest(deleteTestCaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestCase`: ModifyTestCase200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestCase&quot;]
 **deleteTestCaseRequest** | [**DeleteTestCaseRequest**](DeleteTestCaseRequest.md) |  | 

### Return type

[**ModifyTestCase200Response**](ModifyTestCase200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestCaseList

> DescribeTestCaseList200Response DescribeTestCaseList(ctx).Authorization(authorization).Action(action).DescribeTestCaseListRequest(describeTestCaseListRequest).Execute()

测试用例列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestCaseList" // string | Action (default to "DescribeTestCaseList")
	describeTestCaseListRequest := *openapiclient.NewDescribeTestCaseListRequest("ProjectName_example") // DescribeTestCaseListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestCaseList(context.Background()).Authorization(authorization).Action(action).DescribeTestCaseListRequest(describeTestCaseListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestCaseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestCaseList`: DescribeTestCaseList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestCaseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestCaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestCaseList&quot;]
 **describeTestCaseListRequest** | [**DescribeTestCaseListRequest**](DescribeTestCaseListRequest.md) |  | 

### Return type

[**DescribeTestCaseList200Response**](DescribeTestCaseList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestCaseSectionList

> DescribeTestCaseSectionList200Response DescribeTestCaseSectionList(ctx).Authorization(authorization).Action(action).DescribeTestCaseSectionListRequest(describeTestCaseSectionListRequest).Execute()

测试用例分组列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestCaseSectionList" // string | Action (default to "DescribeTestCaseSectionList")
	describeTestCaseSectionListRequest := *openapiclient.NewDescribeTestCaseSectionListRequest("ProjectName_example") // DescribeTestCaseSectionListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestCaseSectionList(context.Background()).Authorization(authorization).Action(action).DescribeTestCaseSectionListRequest(describeTestCaseSectionListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestCaseSectionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestCaseSectionList`: DescribeTestCaseSectionList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestCaseSectionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestCaseSectionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestCaseSectionList&quot;]
 **describeTestCaseSectionListRequest** | [**DescribeTestCaseSectionListRequest**](DescribeTestCaseSectionListRequest.md) |  | 

### Return type

[**DescribeTestCaseSectionList200Response**](DescribeTestCaseSectionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestDefectList

> DescribeTestDefectList200Response DescribeTestDefectList(ctx).Authorization(authorization).Action(action).DescribeTestDefectListRequest(describeTestDefectListRequest).Execute()

测试任务关联的缺陷列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestDefectList" // string | Action (default to "DescribeTestDefectList")
	describeTestDefectListRequest := *openapiclient.NewDescribeTestDefectListRequest("ProjectName_example") // DescribeTestDefectListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestDefectList(context.Background()).Authorization(authorization).Action(action).DescribeTestDefectListRequest(describeTestDefectListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestDefectList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestDefectList`: DescribeTestDefectList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestDefectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestDefectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestDefectList&quot;]
 **describeTestDefectListRequest** | [**DescribeTestDefectListRequest**](DescribeTestDefectListRequest.md) |  | 

### Return type

[**DescribeTestDefectList200Response**](DescribeTestDefectList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestList

> DescribeTestList200Response DescribeTestList(ctx).Authorization(authorization).Action(action).DescribeTestListRequest(describeTestListRequest).Execute()

测试任务列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestList" // string | Action (default to "DescribeTestList")
	describeTestListRequest := *openapiclient.NewDescribeTestListRequest("ProjectName_example") // DescribeTestListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestList(context.Background()).Authorization(authorization).Action(action).DescribeTestListRequest(describeTestListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestList`: DescribeTestList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestList&quot;]
 **describeTestListRequest** | [**DescribeTestListRequest**](DescribeTestListRequest.md) |  | 

### Return type

[**DescribeTestList200Response**](DescribeTestList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestRun

> DescribeTestRun200Response DescribeTestRun(ctx).Authorization(authorization).Action(action).DeleteTestRunRequest(deleteTestRunRequest).Execute()

测试计划详情



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestRun" // string | Action (default to "DescribeTestRun")
	deleteTestRunRequest := *openapiclient.NewDeleteTestRunRequest("ProjectName_example", int32(123)) // DeleteTestRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestRun(context.Background()).Authorization(authorization).Action(action).DeleteTestRunRequest(deleteTestRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestRun`: DescribeTestRun200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestRun&quot;]
 **deleteTestRunRequest** | [**DeleteTestRunRequest**](DeleteTestRunRequest.md) |  | 

### Return type

[**DescribeTestRun200Response**](DescribeTestRun200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTestRunList

> DescribeTestRunList200Response DescribeTestRunList(ctx).Authorization(authorization).Action(action).DescribeTestRunListRequest(describeTestRunListRequest).Execute()

测试计划列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeTestRunList" // string | Action (default to "DescribeTestRunList")
	describeTestRunListRequest := *openapiclient.NewDescribeTestRunListRequest("ProjectName_example") // DescribeTestRunListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeTestRunList(context.Background()).Authorization(authorization).Action(action).DescribeTestRunListRequest(describeTestRunListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeTestRunList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTestRunList`: DescribeTestRunList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeTestRunList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTestRunListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeTestRunList&quot;]
 **describeTestRunListRequest** | [**DescribeTestRunListRequest**](DescribeTestRunListRequest.md) |  | 

### Return type

[**DescribeTestRunList200Response**](DescribeTestRunList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUserGroups

> DescribeUserGroups200Response DescribeUserGroups(ctx).Authorization(authorization).Action(action).DescribeUserGroupsRequest(describeUserGroupsRequest).Execute()

用户组列表分页查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeUserGroups" // string | Action (default to "DescribeUserGroups")
	describeUserGroupsRequest := *openapiclient.NewDescribeUserGroupsRequest(*openapiclient.NewDescribeUserGroupsRequestFilter(), int32(123), int32(123)) // DescribeUserGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeUserGroups(context.Background()).Authorization(authorization).Action(action).DescribeUserGroupsRequest(describeUserGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeUserGroups`: DescribeUserGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeUserGroups&quot;]
 **describeUserGroupsRequest** | [**DescribeUserGroupsRequest**](DescribeUserGroupsRequest.md) |  | 

### Return type

[**DescribeUserGroups200Response**](DescribeUserGroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUserProjects

> DescribeProjectLabels200Response DescribeUserProjects(ctx).Authorization(authorization).Action(action).DescribeUserProjectsRequest(describeUserProjectsRequest).Execute()

项目列表查询（已加入的项目）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeUserProjects" // string | Action (default to "DescribeUserProjects")
	describeUserProjectsRequest := *openapiclient.NewDescribeUserProjectsRequest(int64(123)) // DescribeUserProjectsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeUserProjects(context.Background()).Authorization(authorization).Action(action).DescribeUserProjectsRequest(describeUserProjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeUserProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeUserProjects`: DescribeProjectLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeUserProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUserProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeUserProjects&quot;]
 **describeUserProjectsRequest** | [**DescribeUserProjectsRequest**](DescribeUserProjectsRequest.md) |  | 

### Return type

[**DescribeProjectLabels200Response**](DescribeProjectLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUsersByGroupId

> DescribeUsersByGroupId200Response DescribeUsersByGroupId(ctx).Authorization(authorization).Action(action).DescribeUsersByGroupIdRequest(describeUsersByGroupIdRequest).Execute()

用户列表查询（根据用户组id分页查询）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeUsersByGroupId" // string | Action (default to "DescribeUsersByGroupId")
	describeUsersByGroupIdRequest := *openapiclient.NewDescribeUsersByGroupIdRequest(int32(123), int32(123), int32(123)) // DescribeUsersByGroupIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeUsersByGroupId(context.Background()).Authorization(authorization).Action(action).DescribeUsersByGroupIdRequest(describeUsersByGroupIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeUsersByGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeUsersByGroupId`: DescribeUsersByGroupId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeUsersByGroupId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUsersByGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeUsersByGroupId&quot;]
 **describeUsersByGroupIdRequest** | [**DescribeUsersByGroupIdRequest**](DescribeUsersByGroupIdRequest.md) |  | 

### Return type

[**DescribeUsersByGroupId200Response**](DescribeUsersByGroupId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUsersOnResourceAndGrantObject

> DescribeUsersOnResourceAndGrantObject200Response DescribeUsersOnResourceAndGrantObject(ctx).Authorization(authorization).Action(action).DescribeUsersOnResourceAndGrantObjectRequest(describeUsersOnResourceAndGrantObjectRequest).Execute()

授权用户列表分页查询（指定资源）



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeUsersOnResourceAndGrantObject" // string | Action (default to "DescribeUsersOnResourceAndGrantObject")
	describeUsersOnResourceAndGrantObjectRequest := *openapiclient.NewDescribeUsersOnResourceAndGrantObjectRequest(*openapiclient.NewDescribeUsersOnResourceAndGrantObjectGrantInfo("GrantObjectId_example", "GrantScope_example"), int32(123), int32(123), *openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example")) // DescribeUsersOnResourceAndGrantObjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeUsersOnResourceAndGrantObject(context.Background()).Authorization(authorization).Action(action).DescribeUsersOnResourceAndGrantObjectRequest(describeUsersOnResourceAndGrantObjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeUsersOnResourceAndGrantObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeUsersOnResourceAndGrantObject`: DescribeUsersOnResourceAndGrantObject200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeUsersOnResourceAndGrantObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUsersOnResourceAndGrantObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeUsersOnResourceAndGrantObject&quot;]
 **describeUsersOnResourceAndGrantObjectRequest** | [**DescribeUsersOnResourceAndGrantObjectRequest**](DescribeUsersOnResourceAndGrantObjectRequest.md) |  | 

### Return type

[**DescribeUsersOnResourceAndGrantObject200Response**](DescribeUsersOnResourceAndGrantObject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkItemSalvage

> DescribeWorkItemSalvage200Response DescribeWorkItemSalvage(ctx).Authorization(authorization).Action(action).DescribeWorkItemSalvageRequest(describeWorkItemSalvageRequest).Execute()

事项分解信息查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeWorkItemSalvage" // string | Action (default to "DescribeWorkItemSalvage")
	describeWorkItemSalvageRequest := *openapiclient.NewDescribeWorkItemSalvageRequest(int64(123), "ProjectName_example") // DescribeWorkItemSalvageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeWorkItemSalvage(context.Background()).Authorization(authorization).Action(action).DescribeWorkItemSalvageRequest(describeWorkItemSalvageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeWorkItemSalvage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkItemSalvage`: DescribeWorkItemSalvage200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeWorkItemSalvage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkItemSalvageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeWorkItemSalvage&quot;]
 **describeWorkItemSalvageRequest** | [**DescribeWorkItemSalvageRequest**](DescribeWorkItemSalvageRequest.md) |  | 

### Return type

[**DescribeWorkItemSalvage200Response**](DescribeWorkItemSalvage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkbenchIssueList

> DescribeWorkbenchIssueList200Response DescribeWorkbenchIssueList(ctx).Authorization(authorization).Action(action).DescribeWorkbenchIssueListRequest(describeWorkbenchIssueListRequest).Execute()

用户在团队内的所有代办事项查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeWorkbenchIssueList" // string | Action (default to "DescribeWorkbenchIssueList")
	describeWorkbenchIssueListRequest := *openapiclient.NewDescribeWorkbenchIssueListRequest() // DescribeWorkbenchIssueListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DescribeWorkbenchIssueList(context.Background()).Authorization(authorization).Action(action).DescribeWorkbenchIssueListRequest(describeWorkbenchIssueListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DescribeWorkbenchIssueList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkbenchIssueList`: DescribeWorkbenchIssueList200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DescribeWorkbenchIssueList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkbenchIssueListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeWorkbenchIssueList&quot;]
 **describeWorkbenchIssueListRequest** | [**DescribeWorkbenchIssueListRequest**](DescribeWorkbenchIssueListRequest.md) |  | 

### Return type

[**DescribeWorkbenchIssueList200Response**](DescribeWorkbenchIssueList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachFromResource

> DeleteAPIDoc200Response DetachFromResource(ctx).Authorization(authorization).Action(action).DetachFromResourceRequest(detachFromResourceRequest).Execute()

授权收回，只收回参数指定的授权，已有其它授权不受影响



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DetachFromResource" // string | Action (default to "DetachFromResource")
	detachFromResourceRequest := *openapiclient.NewDetachFromResourceRequest([]openapiclient.GrantInfo{*openapiclient.NewGrantInfo("GrantObjectId_example", "GrantScope_example", int64(123))}, *openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example")) // DetachFromResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DetachFromResource(context.Background()).Authorization(authorization).Action(action).DetachFromResourceRequest(detachFromResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DetachFromResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachFromResource`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DetachFromResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetachFromResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DetachFromResource&quot;]
 **detachFromResourceRequest** | [**DetachFromResourceRequest**](DetachFromResourceRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachResourceScopeOnPolicy

> DeleteAPIDoc200Response DetachResourceScopeOnPolicy(ctx).Authorization(authorization).Action(action).DetachResourceScopeOnPolicyRequest(detachResourceScopeOnPolicyRequest).Execute()

权限组可用资源删除，只删除参数指定的资源，已有的其它资源不受影响



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DetachResourceScopeOnPolicy" // string | Action (default to "DetachResourceScopeOnPolicy")
	detachResourceScopeOnPolicyRequest := *openapiclient.NewDetachResourceScopeOnPolicyRequest(int64(123), []openapiclient.ResourceInfoOfPolicyScope{*openapiclient.NewResourceInfoOfPolicyScope("ResourceId_example", "ResourceType_example")}) // DetachResourceScopeOnPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.DetachResourceScopeOnPolicy(context.Background()).Authorization(authorization).Action(action).DetachResourceScopeOnPolicyRequest(detachResourceScopeOnPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DetachResourceScopeOnPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachResourceScopeOnPolicy`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DetachResourceScopeOnPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetachResourceScopeOnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DetachResourceScopeOnPolicy&quot;]
 **detachResourceScopeOnPolicyRequest** | [**DetachResourceScopeOnPolicyRequest**](DetachResourceScopeOnPolicyRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForbiddenArtifactVersion

> DeleteAPIDoc200Response ForbiddenArtifactVersion(ctx).Authorization(authorization).Action(action).ForbiddenArtifactVersionRequest(forbiddenArtifactVersionRequest).Execute()

制品版本下载 禁止、解禁



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ForbiddenArtifactVersion" // string | Action (default to "ForbiddenArtifactVersion")
	forbiddenArtifactVersionRequest := *openapiclient.NewForbiddenArtifactVersionRequest("ForbiddenAction_example", "Package_example", "PackageVersion_example", int32(123), "Repository_example") // ForbiddenArtifactVersionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ForbiddenArtifactVersion(context.Background()).Authorization(authorization).Action(action).ForbiddenArtifactVersionRequest(forbiddenArtifactVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ForbiddenArtifactVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForbiddenArtifactVersion`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ForbiddenArtifactVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForbiddenArtifactVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ForbiddenArtifactVersion&quot;]
 **forbiddenArtifactVersionRequest** | [**ForbiddenArtifactVersionRequest**](ForbiddenArtifactVersionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyArtifactCredit

> ModifyArtifactCredit200Response ModifyArtifactCredit(ctx).Authorization(authorization).Action(action).ModifyArtifactCreditRequest(modifyArtifactCreditRequest).Execute()

制品授信清单修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyArtifactCredit" // string | Action (default to "ModifyArtifactCredit")
	modifyArtifactCreditRequest := *openapiclient.NewModifyArtifactCreditRequest([]openapiclient.ArtifactsOpenApiCreateArtifactCreditsRangeData{*openapiclient.NewArtifactsOpenApiCreateArtifactCreditsRangeData("RangeType_example", int64(123))}, false, int64(123), []openapiclient.ArtifactsOpenApiArtifactCreditsRuleData{*openapiclient.NewArtifactsOpenApiArtifactCreditsRuleData("Version_example", int32(123), "PkgNameAlgorithm_example", "PkgName_example", "VersionAlgorithm_example")}, "Name_example") // ModifyArtifactCreditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyArtifactCredit(context.Background()).Authorization(authorization).Action(action).ModifyArtifactCreditRequest(modifyArtifactCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyArtifactCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyArtifactCredit`: ModifyArtifactCredit200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyArtifactCredit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyArtifactCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyArtifactCredit&quot;]
 **modifyArtifactCreditRequest** | [**ModifyArtifactCreditRequest**](ModifyArtifactCreditRequest.md) |  | 

### Return type

[**ModifyArtifactCredit200Response**](ModifyArtifactCredit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyArtifactProperties

> DeleteAPIDoc200Response ModifyArtifactProperties(ctx).Authorization(authorization).Action(action).ModifyArtifactPropertiesRequest(modifyArtifactPropertiesRequest).Execute()

制品属性修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyArtifactProperties" // string | Action (default to "ModifyArtifactProperties")
	modifyArtifactPropertiesRequest := *openapiclient.NewModifyArtifactPropertiesRequest("Package_example", "PackageVersion_example", int32(123), []openapiclient.ArtifactPropertyBean{*openapiclient.NewArtifactPropertyBean("Name_example", "Value_example")}, "Repository_example") // ModifyArtifactPropertiesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyArtifactProperties(context.Background()).Authorization(authorization).Action(action).ModifyArtifactPropertiesRequest(modifyArtifactPropertiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyArtifactProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyArtifactProperties`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyArtifactProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyArtifactPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyArtifactProperties&quot;]
 **modifyArtifactPropertiesRequest** | [**ModifyArtifactPropertiesRequest**](ModifyArtifactPropertiesRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBranchProtection

> ModifyBranchProtection200Response ModifyBranchProtection(ctx).Authorization(authorization).Action(action).ModifyBranchProtectionRequest(modifyBranchProtectionRequest).Execute()

仓库设置-修改保护分支规则相关信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyBranchProtection" // string | Action (default to "ModifyBranchProtection")
	modifyBranchProtectionRequest := *openapiclient.NewModifyBranchProtectionRequest(*openapiclient.NewBranchProtection(), int64(123)) // ModifyBranchProtectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyBranchProtection(context.Background()).Authorization(authorization).Action(action).ModifyBranchProtectionRequest(modifyBranchProtectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyBranchProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBranchProtection`: ModifyBranchProtection200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyBranchProtection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyBranchProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyBranchProtection&quot;]
 **modifyBranchProtectionRequest** | [**ModifyBranchProtectionRequest**](ModifyBranchProtectionRequest.md) |  | 

### Return type

[**ModifyBranchProtection200Response**](ModifyBranchProtection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBranchProtectionMemberPermission

> DeleteAPIDoc200Response ModifyBranchProtectionMemberPermission(ctx).Authorization(authorization).Action(action).ModifyBranchProtectionMemberPermissionRequest(modifyBranchProtectionMemberPermissionRequest).Execute()

仓库设置-更改保护分支管理员权限



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyBranchProtectionMemberPermission" // string | Action (default to "ModifyBranchProtectionMemberPermission")
	modifyBranchProtectionMemberPermissionRequest := *openapiclient.NewModifyBranchProtectionMemberPermissionRequest(false, "BranchRuleName_example", "DepotPath_example", "UserGlobalKey_example") // ModifyBranchProtectionMemberPermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyBranchProtectionMemberPermission(context.Background()).Authorization(authorization).Action(action).ModifyBranchProtectionMemberPermissionRequest(modifyBranchProtectionMemberPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyBranchProtectionMemberPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyBranchProtectionMemberPermission`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyBranchProtectionMemberPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyBranchProtectionMemberPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyBranchProtectionMemberPermission&quot;]
 **modifyBranchProtectionMemberPermissionRequest** | [**ModifyBranchProtectionMemberPermissionRequest**](ModifyBranchProtectionMemberPermissionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCdCloudAccount

> ModifyCdCloudAccount200Response ModifyCdCloudAccount(ctx).Authorization(authorization).Action(action).ModifyCdCloudAccountRequest(modifyCdCloudAccountRequest).Execute()

CD 云账号更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyCdCloudAccount" // string | Action (default to "ModifyCdCloudAccount")
	modifyCdCloudAccountRequest := *openapiclient.NewModifyCdCloudAccountRequest(*openapiclient.NewCloudAccountCredential(), int64(123), "Name_example") // ModifyCdCloudAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyCdCloudAccount(context.Background()).Authorization(authorization).Action(action).ModifyCdCloudAccountRequest(modifyCdCloudAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyCdCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCdCloudAccount`: ModifyCdCloudAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyCdCloudAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCdCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyCdCloudAccount&quot;]
 **modifyCdCloudAccountRequest** | [**ModifyCdCloudAccountRequest**](ModifyCdCloudAccountRequest.md) |  | 

### Return type

[**ModifyCdCloudAccount200Response**](ModifyCdCloudAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCdHostServerGroup

> ModifyCdHostServerGroup200Response ModifyCdHostServerGroup(ctx).Authorization(authorization).Action(action).ModifyCdHostServerGroupRequest(modifyCdHostServerGroupRequest).Execute()

CD 主机组修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyCdHostServerGroup" // string | Action (default to "ModifyCdHostServerGroup")
	modifyCdHostServerGroupRequest := *openapiclient.NewModifyCdHostServerGroupRequest(int64(123), "AuthMethod_example", "DisplayName_example", int64(123), []string{"Ips_example"}, int64(123), "UserName_example") // ModifyCdHostServerGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyCdHostServerGroup(context.Background()).Authorization(authorization).Action(action).ModifyCdHostServerGroupRequest(modifyCdHostServerGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyCdHostServerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCdHostServerGroup`: ModifyCdHostServerGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyCdHostServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCdHostServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyCdHostServerGroup&quot;]
 **modifyCdHostServerGroupRequest** | [**ModifyCdHostServerGroupRequest**](ModifyCdHostServerGroupRequest.md) |  | 

### Return type

[**ModifyCdHostServerGroup200Response**](ModifyCdHostServerGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCdPipeline

> ModifyCdPipeline200Response ModifyCdPipeline(ctx).Authorization(authorization).Action(action).ModifyCdPipelineRequest(modifyCdPipelineRequest).Execute()

CD 部署流程修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyCdPipeline" // string | Action (default to "ModifyCdPipeline")
	modifyCdPipelineRequest := *openapiclient.NewModifyCdPipelineRequest("PipelineConfigId_example", "PipelineJsonContent_example") // ModifyCdPipelineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyCdPipeline(context.Background()).Authorization(authorization).Action(action).ModifyCdPipelineRequest(modifyCdPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyCdPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCdPipeline`: ModifyCdPipeline200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyCdPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCdPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyCdPipeline&quot;]
 **modifyCdPipelineRequest** | [**ModifyCdPipelineRequest**](ModifyCdPipelineRequest.md) |  | 

### Return type

[**ModifyCdPipeline200Response**](ModifyCdPipeline200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyChooseDepotSpec

> ModifyChooseDepotSpec200Response ModifyChooseDepotSpec(ctx).Authorization(authorization).Action(action).ModifyChooseDepotSpecRequest(modifyChooseDepotSpecRequest).Execute()

仓库设置-使用、取消使用仓库规范



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyChooseDepotSpec" // string | Action (default to "ModifyChooseDepotSpec")
	modifyChooseDepotSpecRequest := *openapiclient.NewModifyChooseDepotSpecRequest("DepotPath_example", "DepotSpecName_example") // ModifyChooseDepotSpecRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyChooseDepotSpec(context.Background()).Authorization(authorization).Action(action).ModifyChooseDepotSpecRequest(modifyChooseDepotSpecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyChooseDepotSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyChooseDepotSpec`: ModifyChooseDepotSpec200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyChooseDepotSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyChooseDepotSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyChooseDepotSpec&quot;]
 **modifyChooseDepotSpecRequest** | [**ModifyChooseDepotSpecRequest**](ModifyChooseDepotSpecRequest.md) |  | 

### Return type

[**ModifyChooseDepotSpec200Response**](ModifyChooseDepotSpec200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCloseMR

> DeleteAPIDoc200Response ModifyCloseMR(ctx).Authorization(authorization).Action(action).ModifyCloseMRRequest(modifyCloseMRRequest).Execute()

合并请求-关闭合并请求



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyCloseMR" // string | Action (default to "ModifyCloseMR")
	modifyCloseMRRequest := *openapiclient.NewModifyCloseMRRequest(int64(123), int64(123)) // ModifyCloseMRRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyCloseMR(context.Background()).Authorization(authorization).Action(action).ModifyCloseMRRequest(modifyCloseMRRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyCloseMR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCloseMR`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyCloseMR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCloseMRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyCloseMR&quot;]
 **modifyCloseMRRequest** | [**ModifyCloseMRRequest**](ModifyCloseMRRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCodingCIAgentEnable

> DeleteAPIDoc200Response ModifyCodingCIAgentEnable(ctx).Authorization(authorization).Action(action).ModifyCodingCIAgentEnableRequest(modifyCodingCIAgentEnableRequest).Execute()

自定义构建节点启用、禁用



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyCodingCIAgentEnable" // string | Action (default to "ModifyCodingCIAgentEnable")
	modifyCodingCIAgentEnableRequest := *openapiclient.NewModifyCodingCIAgentEnableRequest(int64(123), false, int64(123)) // ModifyCodingCIAgentEnableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyCodingCIAgentEnable(context.Background()).Authorization(authorization).Action(action).ModifyCodingCIAgentEnableRequest(modifyCodingCIAgentEnableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyCodingCIAgentEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCodingCIAgentEnable`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyCodingCIAgentEnable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCodingCIAgentEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyCodingCIAgentEnable&quot;]
 **modifyCodingCIAgentEnableRequest** | [**ModifyCodingCIAgentEnableRequest**](ModifyCodingCIAgentEnableRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCodingCIJob

> DeleteAPIDoc200Response ModifyCodingCIJob(ctx).Authorization(authorization).Action(action).ModifyCodingCIJobRequest(modifyCodingCIJobRequest).Execute()

构建计划修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyCodingCIJob" // string | Action (default to "ModifyCodingCIJob")
	modifyCodingCIJobRequest := *openapiclient.NewModifyCodingCIJobRequest(false, false, int32(123), "DepotType_example", "ExecuteIn_example", "HookType_example", int32(123), "JenkinsFileFromType_example", "JobFromType_example", int32(123), "TriggerRemind_example") // ModifyCodingCIJobRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyCodingCIJob(context.Background()).Authorization(authorization).Action(action).ModifyCodingCIJobRequest(modifyCodingCIJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyCodingCIJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCodingCIJob`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyCodingCIJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCodingCIJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyCodingCIJob&quot;]
 **modifyCodingCIJobRequest** | [**ModifyCodingCIJobRequest**](ModifyCodingCIJobRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDefaultBranch

> DeleteAPIDoc200Response ModifyDefaultBranch(ctx).Authorization(authorization).Action(action).ModifyDefaultBranchRequest(modifyDefaultBranchRequest).Execute()

仓库设置-修改仓库默认分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDefaultBranch" // string | Action (default to "ModifyDefaultBranch")
	modifyDefaultBranchRequest := *openapiclient.NewModifyDefaultBranchRequest("BranchName_example", int64(123), int64(123)) // ModifyDefaultBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDefaultBranch(context.Background()).Authorization(authorization).Action(action).ModifyDefaultBranchRequest(modifyDefaultBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDefaultBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDefaultBranch`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDefaultBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDefaultBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDefaultBranch&quot;]
 **modifyDefaultBranchRequest** | [**ModifyDefaultBranchRequest**](ModifyDefaultBranchRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDefectRelatedRequirement

> DeleteAPIDoc200Response ModifyDefectRelatedRequirement(ctx).Authorization(authorization).Action(action).ModifyDefectRelatedRequirementRequest(modifyDefectRelatedRequirementRequest).Execute()

缺陷所属的需求修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDefectRelatedRequirement" // string | Action (default to "ModifyDefectRelatedRequirement")
	modifyDefectRelatedRequirementRequest := *openapiclient.NewModifyDefectRelatedRequirementRequest(int64(123), "ProjectName_example", int64(123)) // ModifyDefectRelatedRequirementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDefectRelatedRequirement(context.Background()).Authorization(authorization).Action(action).ModifyDefectRelatedRequirementRequest(modifyDefectRelatedRequirementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDefectRelatedRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDefectRelatedRequirement`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDefectRelatedRequirement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDefectRelatedRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDefectRelatedRequirement&quot;]
 **modifyDefectRelatedRequirementRequest** | [**ModifyDefectRelatedRequirementRequest**](ModifyDefectRelatedRequirementRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepartment

> ModifyDepartment200Response ModifyDepartment(ctx).Authorization(authorization).Action(action).ModifyDepartmentRequest(modifyDepartmentRequest).Execute()

部门信息修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepartment" // string | Action (default to "ModifyDepartment")
	modifyDepartmentRequest := *openapiclient.NewModifyDepartmentRequest(int64(123), "Name_example", int64(123)) // ModifyDepartmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepartment(context.Background()).Authorization(authorization).Action(action).ModifyDepartmentRequest(modifyDepartmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepartment`: ModifyDepartment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepartment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepartment&quot;]
 **modifyDepartmentRequest** | [**ModifyDepartmentRequest**](ModifyDepartmentRequest.md) |  | 

### Return type

[**ModifyDepartment200Response**](ModifyDepartment200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepartmentAssignee

> DeleteAPIDoc200Response ModifyDepartmentAssignee(ctx).Authorization(authorization).Action(action).ModifyDepartmentAssigneeRequest(modifyDepartmentAssigneeRequest).Execute()

部门负责人管理



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepartmentAssignee" // string | Action (default to "ModifyDepartmentAssignee")
	modifyDepartmentAssigneeRequest := *openapiclient.NewModifyDepartmentAssigneeRequest(int64(123), []int64{int64(123)}) // ModifyDepartmentAssigneeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepartmentAssignee(context.Background()).Authorization(authorization).Action(action).ModifyDepartmentAssigneeRequest(modifyDepartmentAssigneeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepartmentAssignee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepartmentAssignee`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepartmentAssignee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepartmentAssigneeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepartmentAssignee&quot;]
 **modifyDepartmentAssigneeRequest** | [**ModifyDepartmentAssigneeRequest**](ModifyDepartmentAssigneeRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepartmentMember

> DeleteAPIDoc200Response ModifyDepartmentMember(ctx).Authorization(authorization).Action(action).ModifyDepartmentMemberRequest(modifyDepartmentMemberRequest).Execute()

部门成员管理



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepartmentMember" // string | Action (default to "ModifyDepartmentMember")
	modifyDepartmentMemberRequest := *openapiclient.NewModifyDepartmentMemberRequest([]int64{int64(123)}, []int64{int64(123)}) // ModifyDepartmentMemberRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepartmentMember(context.Background()).Authorization(authorization).Action(action).ModifyDepartmentMemberRequest(modifyDepartmentMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepartmentMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepartmentMember`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepartmentMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepartmentMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepartmentMember&quot;]
 **modifyDepartmentMemberRequest** | [**ModifyDepartmentMemberRequest**](ModifyDepartmentMemberRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotDescription

> DeleteAPIDoc200Response ModifyDepotDescription(ctx).Authorization(authorization).Action(action).ModifyDepotDescriptionRequest(modifyDepotDescriptionRequest).Execute()

仓库信息-修改仓库描述



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotDescription" // string | Action (default to "ModifyDepotDescription")
	modifyDepotDescriptionRequest := *openapiclient.NewModifyDepotDescriptionRequest(int64(123), "Description_example", int64(123)) // ModifyDepotDescriptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotDescription(context.Background()).Authorization(authorization).Action(action).ModifyDepotDescriptionRequest(modifyDepotDescriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotDescription`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotDescription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotDescription&quot;]
 **modifyDepotDescriptionRequest** | [**ModifyDepotDescriptionRequest**](ModifyDepotDescriptionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotFilePushRule

> ModifyDepotFilePushRule200Response ModifyDepotFilePushRule(ctx).Authorization(authorization).Action(action).ModifyDepotFilePushRuleRequest(modifyDepotFilePushRuleRequest).Execute()

仓库设置-修改git仓库文件推送规则



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotFilePushRule" // string | Action (default to "ModifyDepotFilePushRule")
	modifyDepotFilePushRuleRequest := *openapiclient.NewModifyDepotFilePushRuleRequest("DepotPath_example", int64(123), false) // ModifyDepotFilePushRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotFilePushRule(context.Background()).Authorization(authorization).Action(action).ModifyDepotFilePushRuleRequest(modifyDepotFilePushRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotFilePushRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotFilePushRule`: ModifyDepotFilePushRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotFilePushRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotFilePushRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotFilePushRule&quot;]
 **modifyDepotFilePushRuleRequest** | [**ModifyDepotFilePushRuleRequest**](ModifyDepotFilePushRuleRequest.md) |  | 

### Return type

[**ModifyDepotFilePushRule200Response**](ModifyDepotFilePushRule200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotFilePushRuleDenyPrivilege

> ModifyDepotFilePushRuleDenyPrivilege200Response ModifyDepotFilePushRuleDenyPrivilege(ctx).Authorization(authorization).Action(action).ModifyDepotFilePushRuleDenyPrivilegeRequest(modifyDepotFilePushRuleDenyPrivilegeRequest).Execute()

仓库设置-修改 git 仓库特权者文件推送权限



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotFilePushRuleDenyPrivilege" // string | Action (default to "ModifyDepotFilePushRuleDenyPrivilege")
	modifyDepotFilePushRuleDenyPrivilegeRequest := *openapiclient.NewModifyDepotFilePushRuleDenyPrivilegeRequest("DepotPath_example", int64(123), false, false, false, "UserGlobalKey_example") // ModifyDepotFilePushRuleDenyPrivilegeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotFilePushRuleDenyPrivilege(context.Background()).Authorization(authorization).Action(action).ModifyDepotFilePushRuleDenyPrivilegeRequest(modifyDepotFilePushRuleDenyPrivilegeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotFilePushRuleDenyPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotFilePushRuleDenyPrivilege`: ModifyDepotFilePushRuleDenyPrivilege200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotFilePushRuleDenyPrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotFilePushRuleDenyPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotFilePushRuleDenyPrivilege&quot;]
 **modifyDepotFilePushRuleDenyPrivilegeRequest** | [**ModifyDepotFilePushRuleDenyPrivilegeRequest**](ModifyDepotFilePushRuleDenyPrivilegeRequest.md) |  | 

### Return type

[**ModifyDepotFilePushRuleDenyPrivilege200Response**](ModifyDepotFilePushRuleDenyPrivilege200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotLevelDepotSpec

> ModifyDepotLevelDepotSpec200Response ModifyDepotLevelDepotSpec(ctx).Authorization(authorization).Action(action).ModifyDepotLevelDepotSpecRequest(modifyDepotLevelDepotSpecRequest).Execute()

仓库设置-修改、新增仓库级别的仓库规范



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotLevelDepotSpec" // string | Action (default to "ModifyDepotLevelDepotSpec")
	modifyDepotLevelDepotSpecRequest := *openapiclient.NewModifyDepotLevelDepotSpecRequest("DepotPath_example", *openapiclient.NewDepotSpecDepotLevelParam()) // ModifyDepotLevelDepotSpecRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotLevelDepotSpec(context.Background()).Authorization(authorization).Action(action).ModifyDepotLevelDepotSpecRequest(modifyDepotLevelDepotSpecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotLevelDepotSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotLevelDepotSpec`: ModifyDepotLevelDepotSpec200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotLevelDepotSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotLevelDepotSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotLevelDepotSpec&quot;]
 **modifyDepotLevelDepotSpecRequest** | [**ModifyDepotLevelDepotSpecRequest**](ModifyDepotLevelDepotSpecRequest.md) |  | 

### Return type

[**ModifyDepotLevelDepotSpec200Response**](ModifyDepotLevelDepotSpec200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotName

> ModifyDepotName200Response ModifyDepotName(ctx).Authorization(authorization).Action(action).ModifyDepotNameRequest(modifyDepotNameRequest).Execute()

仓库信息-修改仓库名称



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotName" // string | Action (default to "ModifyDepotName")
	modifyDepotNameRequest := *openapiclient.NewModifyDepotNameRequest(int64(123), "DepotName_example", int64(123)) // ModifyDepotNameRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotName(context.Background()).Authorization(authorization).Action(action).ModifyDepotNameRequest(modifyDepotNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotName`: ModifyDepotName200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotName&quot;]
 **modifyDepotNameRequest** | [**ModifyDepotNameRequest**](ModifyDepotNameRequest.md) |  | 

### Return type

[**ModifyDepotName200Response**](ModifyDepotName200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotPushSetting

> ModifyDepotPushSetting200Response ModifyDepotPushSetting(ctx).Authorization(authorization).Action(action).ModifyDepotPushSettingRequest(modifyDepotPushSettingRequest).Execute()

仓库设置-修改仓库推送设置



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotPushSetting" // string | Action (default to "ModifyDepotPushSetting")
	modifyDepotPushSettingRequest := *openapiclient.NewModifyDepotPushSettingRequest("DepotPath_example", *openapiclient.NewDepotPushSetting()) // ModifyDepotPushSettingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotPushSetting(context.Background()).Authorization(authorization).Action(action).ModifyDepotPushSettingRequest(modifyDepotPushSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotPushSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotPushSetting`: ModifyDepotPushSetting200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotPushSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotPushSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotPushSetting&quot;]
 **modifyDepotPushSettingRequest** | [**ModifyDepotPushSettingRequest**](ModifyDepotPushSettingRequest.md) |  | 

### Return type

[**ModifyDepotPushSetting200Response**](ModifyDepotPushSetting200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotQuota

> DeleteAPIDoc200Response ModifyDepotQuota(ctx).Authorization(authorization).Action(action).ModifyDepotQuotaRequest(modifyDepotQuotaRequest).Execute()

仓库信息-修改仓库容量



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotQuota" // string | Action (default to "ModifyDepotQuota")
	modifyDepotQuotaRequest := *openapiclient.NewModifyDepotQuotaRequest("DepotPath_example", "DepotQuotaSize_example") // ModifyDepotQuotaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotQuota(context.Background()).Authorization(authorization).Action(action).ModifyDepotQuotaRequest(modifyDepotQuotaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotQuota`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotQuota&quot;]
 **modifyDepotQuotaRequest** | [**ModifyDepotQuotaRequest**](ModifyDepotQuotaRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotSettings

> DeleteAPIDoc200Response ModifyDepotSettings(ctx).Authorization(authorization).Action(action).ModifyDepotSettingsRequest(modifyDepotSettingsRequest).Execute()

仓库设置-修改仓库设置



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotSettings" // string | Action (default to "ModifyDepotSettings")
	modifyDepotSettingsRequest := *openapiclient.NewModifyDepotSettingsRequest("DepotPath_example", false) // ModifyDepotSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotSettings(context.Background()).Authorization(authorization).Action(action).ModifyDepotSettingsRequest(modifyDepotSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotSettings`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotSettings&quot;]
 **modifyDepotSettingsRequest** | [**ModifyDepotSettingsRequest**](ModifyDepotSettingsRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDepotSharedSetting

> DeleteAPIDoc200Response ModifyDepotSharedSetting(ctx).Authorization(authorization).Action(action).ModifyDepotSharedSettingRequest(modifyDepotSharedSettingRequest).Execute()

仓库设置-修改仓库是否开源设置



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyDepotSharedSetting" // string | Action (default to "ModifyDepotSharedSetting")
	modifyDepotSharedSettingRequest := *openapiclient.NewModifyDepotSharedSettingRequest("DepotPath_example", false) // ModifyDepotSharedSettingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyDepotSharedSetting(context.Background()).Authorization(authorization).Action(action).ModifyDepotSharedSettingRequest(modifyDepotSharedSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyDepotSharedSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDepotSharedSetting`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyDepotSharedSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyDepotSharedSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyDepotSharedSetting&quot;]
 **modifyDepotSharedSettingRequest** | [**ModifyDepotSharedSettingRequest**](ModifyDepotSharedSettingRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitCherryPick

> DeleteAPIDoc200Response ModifyGitCherryPick(ctx).Authorization(authorization).Action(action).ModifyGitCherryPickRequest(modifyGitCherryPickRequest).Execute()

Git提交-将某次提交cherry-pick到指定分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitCherryPick" // string | Action (default to "ModifyGitCherryPick")
	modifyGitCherryPickRequest := *openapiclient.NewModifyGitCherryPickRequest("BranchName_example", int64(123), "Sha_example") // ModifyGitCherryPickRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitCherryPick(context.Background()).Authorization(authorization).Action(action).ModifyGitCherryPickRequest(modifyGitCherryPickRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitCherryPick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitCherryPick`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitCherryPick`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitCherryPickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitCherryPick&quot;]
 **modifyGitCherryPickRequest** | [**ModifyGitCherryPickRequest**](ModifyGitCherryPickRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitCommitRevert

> DeleteAPIDoc200Response ModifyGitCommitRevert(ctx).Authorization(authorization).Action(action).ModifyGitCommitRevertRequest(modifyGitCommitRevertRequest).Execute()

Git提交-还原某次提交



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitCommitRevert" // string | Action (default to "ModifyGitCommitRevert")
	modifyGitCommitRevertRequest := *openapiclient.NewModifyGitCommitRevertRequest("BranchName_example", int64(123), "Message_example", "Sha_example") // ModifyGitCommitRevertRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitCommitRevert(context.Background()).Authorization(authorization).Action(action).ModifyGitCommitRevertRequest(modifyGitCommitRevertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitCommitRevert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitCommitRevert`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitCommitRevert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitCommitRevertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitCommitRevert&quot;]
 **modifyGitCommitRevertRequest** | [**ModifyGitCommitRevertRequest**](ModifyGitCommitRevertRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitCommitStatus

> DeleteAPIDoc200Response ModifyGitCommitStatus(ctx).Authorization(authorization).Action(action).ModifyGitCommitStatusRequest(modifyGitCommitStatusRequest).Execute()

Git提交-修改提交对应的流水线状态



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitCommitStatus" // string | Action (default to "ModifyGitCommitStatus")
	modifyGitCommitStatusRequest := *openapiclient.NewModifyGitCommitStatusRequest("CommitSha_example", "Context_example", "Description_example", "State_example", "TargetURL_example") // ModifyGitCommitStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitCommitStatus(context.Background()).Authorization(authorization).Action(action).ModifyGitCommitStatusRequest(modifyGitCommitStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitCommitStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitCommitStatus`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitCommitStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitCommitStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitCommitStatus&quot;]
 **modifyGitCommitStatusRequest** | [**ModifyGitCommitStatusRequest**](ModifyGitCommitStatusRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitDepotArchive

> DeleteAPIDoc200Response ModifyGitDepotArchive(ctx).Authorization(authorization).Action(action).ModifyGitDepotArchiveRequest(modifyGitDepotArchiveRequest).Execute()

仓库设置-仓库归档



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitDepotArchive" // string | Action (default to "ModifyGitDepotArchive")
	modifyGitDepotArchiveRequest := *openapiclient.NewModifyGitDepotArchiveRequest(int64(123)) // ModifyGitDepotArchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitDepotArchive(context.Background()).Authorization(authorization).Action(action).ModifyGitDepotArchiveRequest(modifyGitDepotArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitDepotArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitDepotArchive`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitDepotArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitDepotArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitDepotArchive&quot;]
 **modifyGitDepotArchiveRequest** | [**ModifyGitDepotArchiveRequest**](ModifyGitDepotArchiveRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitDepotUnarchive

> DeleteAPIDoc200Response ModifyGitDepotUnarchive(ctx).Authorization(authorization).Action(action).ModifyGitDepotUnarchiveRequest(modifyGitDepotUnarchiveRequest).Execute()

仓库设置-仓库解除归档



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitDepotUnarchive" // string | Action (default to "ModifyGitDepotUnarchive")
	modifyGitDepotUnarchiveRequest := *openapiclient.NewModifyGitDepotUnarchiveRequest("DepotId_example") // ModifyGitDepotUnarchiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitDepotUnarchive(context.Background()).Authorization(authorization).Action(action).ModifyGitDepotUnarchiveRequest(modifyGitDepotUnarchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitDepotUnarchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitDepotUnarchive`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitDepotUnarchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitDepotUnarchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitDepotUnarchive&quot;]
 **modifyGitDepotUnarchiveRequest** | [**ModifyGitDepotUnarchiveRequest**](ModifyGitDepotUnarchiveRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitFiles

> ModifyGitFiles200Response ModifyGitFiles(ctx).Authorization(authorization).Action(action).ModifyGitFilesRequest(modifyGitFilesRequest).Execute()

Git提交-修改仓库某文件



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitFiles" // string | Action (default to "ModifyGitFiles")
	modifyGitFilesRequest := *openapiclient.NewModifyGitFilesRequest(int64(123), []openapiclient.GitFile{*openapiclient.NewGitFile("Content_example", "Path_example")}, "LastCommitSha_example", "Message_example", "Ref_example") // ModifyGitFilesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitFiles(context.Background()).Authorization(authorization).Action(action).ModifyGitFilesRequest(modifyGitFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitFiles`: ModifyGitFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitFiles&quot;]
 **modifyGitFilesRequest** | [**ModifyGitFilesRequest**](ModifyGitFilesRequest.md) |  | 

### Return type

[**ModifyGitFiles200Response**](ModifyGitFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitMergeBranch

> ModifyGitMergeBranch200Response ModifyGitMergeBranch(ctx).Authorization(authorization).Action(action).ModifyGitMergeBranchRequest(modifyGitMergeBranchRequest).Execute()

合并请求-将源分支的改动合并到目标分支



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitMergeBranch" // string | Action (default to "ModifyGitMergeBranch")
	modifyGitMergeBranchRequest := *openapiclient.NewModifyGitMergeBranchRequest("CommitMessage_example", "DepotPath_example", "FromBranch_example", "ToBranch_example") // ModifyGitMergeBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitMergeBranch(context.Background()).Authorization(authorization).Action(action).ModifyGitMergeBranchRequest(modifyGitMergeBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitMergeBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitMergeBranch`: ModifyGitMergeBranch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitMergeBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitMergeBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitMergeBranch&quot;]
 **modifyGitMergeBranchRequest** | [**ModifyGitMergeBranchRequest**](ModifyGitMergeBranchRequest.md) |  | 

### Return type

[**ModifyGitMergeBranch200Response**](ModifyGitMergeBranch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitMergeRequest

> ModifyGitMergeRequest200Response ModifyGitMergeRequest(ctx).Authorization(authorization).Action(action).ModifyGitMergeRequestRequest(modifyGitMergeRequestRequest).Execute()

合并请求-修改合并请求的标题和描述信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitMergeRequest" // string | Action (default to "ModifyGitMergeRequest")
	modifyGitMergeRequestRequest := *openapiclient.NewModifyGitMergeRequestRequest(int64(123), int64(123)) // ModifyGitMergeRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitMergeRequest(context.Background()).Authorization(authorization).Action(action).ModifyGitMergeRequestRequest(modifyGitMergeRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitMergeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitMergeRequest`: ModifyGitMergeRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitMergeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitMergeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitMergeRequest&quot;]
 **modifyGitMergeRequestRequest** | [**ModifyGitMergeRequestRequest**](ModifyGitMergeRequestRequest.md) |  | 

### Return type

[**ModifyGitMergeRequest200Response**](ModifyGitMergeRequest200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitMergeRequestRebase

> DeleteAPIDoc200Response ModifyGitMergeRequestRebase(ctx).Authorization(authorization).Action(action).ModifyGitMergeRequestRebaseRequest(modifyGitMergeRequestRebaseRequest).Execute()

合并请求-合并请求中的源分支进行rebase操作



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitMergeRequestRebase" // string | Action (default to "ModifyGitMergeRequestRebase")
	modifyGitMergeRequestRebaseRequest := *openapiclient.NewModifyGitMergeRequestRebaseRequest(int64(123), int64(123)) // ModifyGitMergeRequestRebaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitMergeRequestRebase(context.Background()).Authorization(authorization).Action(action).ModifyGitMergeRequestRebaseRequest(modifyGitMergeRequestRebaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitMergeRequestRebase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitMergeRequestRebase`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitMergeRequestRebase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitMergeRequestRebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitMergeRequestRebase&quot;]
 **modifyGitMergeRequestRebaseRequest** | [**ModifyGitMergeRequestRebaseRequest**](ModifyGitMergeRequestRebaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitRebase

> DeleteAPIDoc200Response ModifyGitRebase(ctx).Authorization(authorization).Action(action).ModifyGitRebaseRequest(modifyGitRebaseRequest).Execute()

仓库信息-git变基操作



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitRebase" // string | Action (default to "ModifyGitRebase")
	modifyGitRebaseRequest := *openapiclient.NewModifyGitRebaseRequest("BaseBranchName_example", int64(123), "SrcBranchName_example") // ModifyGitRebaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitRebase(context.Background()).Authorization(authorization).Action(action).ModifyGitRebaseRequest(modifyGitRebaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitRebase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitRebase`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitRebase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitRebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitRebase&quot;]
 **modifyGitRebaseRequest** | [**ModifyGitRebaseRequest**](ModifyGitRebaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitRelease

> DeleteAPIDoc200Response ModifyGitRelease(ctx).Authorization(authorization).Action(action).ModifyGitReleaseRequest(modifyGitReleaseRequest).Execute()

版本信息-修改仓库版本信息



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitRelease" // string | Action (default to "ModifyGitRelease")
	modifyGitReleaseRequest := *openapiclient.NewModifyGitReleaseRequest(int64(123), int64(123), "TagName_example") // ModifyGitReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitRelease(context.Background()).Authorization(authorization).Action(action).ModifyGitReleaseRequest(modifyGitReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitRelease`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitRelease&quot;]
 **modifyGitReleaseRequest** | [**ModifyGitReleaseRequest**](ModifyGitReleaseRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyGitTransfer

> ModifyGitTransfer200Response ModifyGitTransfer(ctx).Authorization(authorization).Action(action).ModifyGitTransferRequest(modifyGitTransferRequest).Execute()

仓库信息-仓库转移至同团队下的其他项目中



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyGitTransfer" // string | Action (default to "ModifyGitTransfer")
	modifyGitTransferRequest := *openapiclient.NewModifyGitTransferRequest(int64(123), int64(123)) // ModifyGitTransferRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyGitTransfer(context.Background()).Authorization(authorization).Action(action).ModifyGitTransferRequest(modifyGitTransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyGitTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyGitTransfer`: ModifyGitTransfer200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyGitTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyGitTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyGitTransfer&quot;]
 **modifyGitTransferRequest** | [**ModifyGitTransferRequest**](ModifyGitTransferRequest.md) |  | 

### Return type

[**ModifyGitTransfer200Response**](ModifyGitTransfer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIssue

> ModifyIssue200Response ModifyIssue(ctx).Authorization(authorization).Action(action).ModifyIssueRequest(modifyIssueRequest).Execute()

事项修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyIssue" // string | Action (default to "ModifyIssue")
	modifyIssueRequest := *openapiclient.NewModifyIssueRequest(int64(123), "ProjectName_example") // ModifyIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyIssue(context.Background()).Authorization(authorization).Action(action).ModifyIssueRequest(modifyIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIssue`: ModifyIssue200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyIssue&quot;]
 **modifyIssueRequest** | [**ModifyIssueRequest**](ModifyIssueRequest.md) |  | 

### Return type

[**ModifyIssue200Response**](ModifyIssue200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIssueComment

> DeleteAPIDoc200Response ModifyIssueComment(ctx).Authorization(authorization).Action(action).ModifyIssueCommentRequest(modifyIssueCommentRequest).Execute()

事项评论修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyIssueComment" // string | Action (default to "ModifyIssueComment")
	modifyIssueCommentRequest := *openapiclient.NewModifyIssueCommentRequest(int64(123), "Content_example", int64(123), "ProjectName_example") // ModifyIssueCommentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyIssueComment(context.Background()).Authorization(authorization).Action(action).ModifyIssueCommentRequest(modifyIssueCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyIssueComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIssueComment`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyIssueComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyIssueCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyIssueComment&quot;]
 **modifyIssueCommentRequest** | [**ModifyIssueCommentRequest**](ModifyIssueCommentRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIssueDescription

> DeleteAPIDoc200Response ModifyIssueDescription(ctx).Authorization(authorization).Action(action).ModifyIssueDescriptionRequest(modifyIssueDescriptionRequest).Execute()

事项描述修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyIssueDescription" // string | Action (default to "ModifyIssueDescription")
	modifyIssueDescriptionRequest := *openapiclient.NewModifyIssueDescriptionRequest("Description_example", int64(123), "ProjectName_example") // ModifyIssueDescriptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyIssueDescription(context.Background()).Authorization(authorization).Action(action).ModifyIssueDescriptionRequest(modifyIssueDescriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyIssueDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIssueDescription`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyIssueDescription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyIssueDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyIssueDescription&quot;]
 **modifyIssueDescriptionRequest** | [**ModifyIssueDescriptionRequest**](ModifyIssueDescriptionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIssueParentRequirement

> DeleteAPIDoc200Response ModifyIssueParentRequirement(ctx).Authorization(authorization).Action(action).ModifyIssueParentRequirementRequest(modifyIssueParentRequirementRequest).Execute()

事项父需求修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyIssueParentRequirement" // string | Action (default to "ModifyIssueParentRequirement")
	modifyIssueParentRequirementRequest := *openapiclient.NewModifyIssueParentRequirementRequest(int64(123), int64(123), "ProjectName_example") // ModifyIssueParentRequirementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyIssueParentRequirement(context.Background()).Authorization(authorization).Action(action).ModifyIssueParentRequirementRequest(modifyIssueParentRequirementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyIssueParentRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIssueParentRequirement`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyIssueParentRequirement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyIssueParentRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyIssueParentRequirement&quot;]
 **modifyIssueParentRequirementRequest** | [**ModifyIssueParentRequirementRequest**](ModifyIssueParentRequirementRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIteration

> ModifyIteration200Response ModifyIteration(ctx).Authorization(authorization).Action(action).ModifyIterationRequest(modifyIterationRequest).Execute()

迭代修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyIteration" // string | Action (default to "ModifyIteration")
	modifyIterationRequest := *openapiclient.NewModifyIterationRequest(int64(123), "Name_example", "ProjectName_example") // ModifyIterationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyIteration(context.Background()).Authorization(authorization).Action(action).ModifyIterationRequest(modifyIterationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyIteration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIteration`: ModifyIteration200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyIteration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyIteration&quot;]
 **modifyIterationRequest** | [**ModifyIterationRequest**](ModifyIterationRequest.md) |  | 

### Return type

[**ModifyIteration200Response**](ModifyIteration200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMergeMR

> ModifyMergeMR200Response ModifyMergeMR(ctx).Authorization(authorization).Action(action).ModifyMergeMRRequest(modifyMergeMRRequest).Execute()

合并信息-执行合并



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyMergeMR" // string | Action (default to "ModifyMergeMR")
	modifyMergeMRRequest := *openapiclient.NewModifyMergeMRRequest(int64(123), false, false, int64(123), "Message_example", false) // ModifyMergeMRRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyMergeMR(context.Background()).Authorization(authorization).Action(action).ModifyMergeMRRequest(modifyMergeMRRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyMergeMR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMergeMR`: ModifyMergeMR200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyMergeMR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyMergeMRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyMergeMR&quot;]
 **modifyMergeMRRequest** | [**ModifyMergeMRRequest**](ModifyMergeMRRequest.md) |  | 

### Return type

[**ModifyMergeMR200Response**](ModifyMergeMR200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMergeRequestBasicSettings

> DeleteAPIDoc200Response ModifyMergeRequestBasicSettings(ctx).Authorization(authorization).Action(action).ModifyMergeRequestBasicSettingsRequest(modifyMergeRequestBasicSettingsRequest).Execute()

仓库设置-修改合并请求基础设置



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyMergeRequestBasicSettings" // string | Action (default to "ModifyMergeRequestBasicSettings")
	modifyMergeRequestBasicSettingsRequest := *openapiclient.NewModifyMergeRequestBasicSettingsRequest(false, false, false, "DepotPath_example", false, false) // ModifyMergeRequestBasicSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyMergeRequestBasicSettings(context.Background()).Authorization(authorization).Action(action).ModifyMergeRequestBasicSettingsRequest(modifyMergeRequestBasicSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyMergeRequestBasicSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMergeRequestBasicSettings`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyMergeRequestBasicSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyMergeRequestBasicSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyMergeRequestBasicSettings&quot;]
 **modifyMergeRequestBasicSettingsRequest** | [**ModifyMergeRequestBasicSettingsRequest**](ModifyMergeRequestBasicSettingsRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMergeRequestMergeCommitMessageTemplate

> DeleteAPIDoc200Response ModifyMergeRequestMergeCommitMessageTemplate(ctx).Authorization(authorization).Action(action).ModifyMergeRequestSquashCommitMessageTemplateRequest(modifyMergeRequestSquashCommitMessageTemplateRequest).Execute()

仓库设置-修改合并请求合并提交消息模版



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyMergeRequestMergeCommitMessageTemplate" // string | Action (default to "ModifyMergeRequestMergeCommitMessageTemplate")
	modifyMergeRequestSquashCommitMessageTemplateRequest := *openapiclient.NewModifyMergeRequestSquashCommitMessageTemplateRequest("CommitMessageTemplate_example", "DepotPath_example") // ModifyMergeRequestSquashCommitMessageTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyMergeRequestMergeCommitMessageTemplate(context.Background()).Authorization(authorization).Action(action).ModifyMergeRequestSquashCommitMessageTemplateRequest(modifyMergeRequestSquashCommitMessageTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyMergeRequestMergeCommitMessageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMergeRequestMergeCommitMessageTemplate`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyMergeRequestMergeCommitMessageTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyMergeRequestMergeCommitMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyMergeRequestMergeCommitMessageTemplate&quot;]
 **modifyMergeRequestSquashCommitMessageTemplateRequest** | [**ModifyMergeRequestSquashCommitMessageTemplateRequest**](ModifyMergeRequestSquashCommitMessageTemplateRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMergeRequestSquashCommitMessageTemplate

> DeleteAPIDoc200Response ModifyMergeRequestSquashCommitMessageTemplate(ctx).Authorization(authorization).Action(action).ModifyMergeRequestSquashCommitMessageTemplateRequest(modifyMergeRequestSquashCommitMessageTemplateRequest).Execute()

仓库设置-修改合并请求合并压缩提交消息模版



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyMergeRequestSquashCommitMessageTemplate" // string | Action (default to "ModifyMergeRequestSquashCommitMessageTemplate")
	modifyMergeRequestSquashCommitMessageTemplateRequest := *openapiclient.NewModifyMergeRequestSquashCommitMessageTemplateRequest("CommitMessageTemplate_example", "DepotPath_example") // ModifyMergeRequestSquashCommitMessageTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyMergeRequestSquashCommitMessageTemplate(context.Background()).Authorization(authorization).Action(action).ModifyMergeRequestSquashCommitMessageTemplateRequest(modifyMergeRequestSquashCommitMessageTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyMergeRequestSquashCommitMessageTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMergeRequestSquashCommitMessageTemplate`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyMergeRequestSquashCommitMessageTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyMergeRequestSquashCommitMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyMergeRequestSquashCommitMessageTemplate&quot;]
 **modifyMergeRequestSquashCommitMessageTemplateRequest** | [**ModifyMergeRequestSquashCommitMessageTemplateRequest**](ModifyMergeRequestSquashCommitMessageTemplateRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPolicy

> ModifyPolicy200Response ModifyPolicy(ctx).Authorization(authorization).Action(action).ModifyPolicyRequest(modifyPolicyRequest).Execute()

权限组修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyPolicy" // string | Action (default to "ModifyPolicy")
	modifyPolicyRequest := *openapiclient.NewModifyPolicyRequest(int64(123), *openapiclient.NewPolicyDocument([]openapiclient.PolicyStatement{*openapiclient.NewPolicyStatement([]string{"Action_example"}, "Effect_example", []string{"Resource_example"})}), []string{"ResourceType_example"}) // ModifyPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyPolicy(context.Background()).Authorization(authorization).Action(action).ModifyPolicyRequest(modifyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPolicy`: ModifyPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyPolicy&quot;]
 **modifyPolicyRequest** | [**ModifyPolicyRequest**](ModifyPolicyRequest.md) |  | 

### Return type

[**ModifyPolicy200Response**](ModifyPolicy200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProject

> DeleteAPIDoc200Response ModifyProject(ctx).Authorization(authorization).Action(action).ModifyProjectRequest(modifyProjectRequest).Execute()

项目信息修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyProject" // string | Action (default to "ModifyProject")
	modifyProjectRequest := *openapiclient.NewModifyProjectRequest("DisplayName_example", "Name_example", int64(123)) // ModifyProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyProject(context.Background()).Authorization(authorization).Action(action).ModifyProjectRequest(modifyProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProject`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyProject&quot;]
 **modifyProjectRequest** | [**ModifyProjectRequest**](ModifyProjectRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProjectAnnouncement

> ModifyProjectAnnouncement200Response ModifyProjectAnnouncement(ctx).Authorization(authorization).Action(action).ModifyProjectAnnouncementRequest(modifyProjectAnnouncementRequest).Execute()

项目公告更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyProjectAnnouncement" // string | Action (default to "ModifyProjectAnnouncement")
	modifyProjectAnnouncementRequest := *openapiclient.NewModifyProjectAnnouncementRequest("ProjectName_example", "Content_example", int64(123)) // ModifyProjectAnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyProjectAnnouncement(context.Background()).Authorization(authorization).Action(action).ModifyProjectAnnouncementRequest(modifyProjectAnnouncementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyProjectAnnouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProjectAnnouncement`: ModifyProjectAnnouncement200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyProjectAnnouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyProjectAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyProjectAnnouncement&quot;]
 **modifyProjectAnnouncementRequest** | [**ModifyProjectAnnouncementRequest**](ModifyProjectAnnouncementRequest.md) |  | 

### Return type

[**ModifyProjectAnnouncement200Response**](ModifyProjectAnnouncement200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProjectLabel

> ModifyProjectLabel200Response ModifyProjectLabel(ctx).Authorization(authorization).Action(action).ModifyProjectLabelRequest(modifyProjectLabelRequest).Execute()

项目标签修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyProjectLabel" // string | Action (default to "ModifyProjectLabel")
	modifyProjectLabelRequest := *openapiclient.NewModifyProjectLabelRequest(int64(123), int64(123), "Name_example", "Color_example") // ModifyProjectLabelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyProjectLabel(context.Background()).Authorization(authorization).Action(action).ModifyProjectLabelRequest(modifyProjectLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyProjectLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProjectLabel`: ModifyProjectLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyProjectLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyProjectLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyProjectLabel&quot;]
 **modifyProjectLabelRequest** | [**ModifyProjectLabelRequest**](ModifyProjectLabelRequest.md) |  | 

### Return type

[**ModifyProjectLabel200Response**](ModifyProjectLabel200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProjectPermission

> DeleteAPIDoc200Response ModifyProjectPermission(ctx).Authorization(authorization).Action(action).ModifyProjectPermissionRequest(modifyProjectPermissionRequest).Execute()

项目成员权限配置(老权限)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyProjectPermission" // string | Action (default to "ModifyProjectPermission")
	modifyProjectPermissionRequest := *openapiclient.NewModifyProjectPermissionRequest(false, int32(123), int32(123), "UserGK_example") // ModifyProjectPermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyProjectPermission(context.Background()).Authorization(authorization).Action(action).ModifyProjectPermissionRequest(modifyProjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyProjectPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyProjectPermission`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyProjectPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyProjectPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyProjectPermission&quot;]
 **modifyProjectPermissionRequest** | [**ModifyProjectPermissionRequest**](ModifyProjectPermissionRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRelease

> ModifyRelease200Response ModifyRelease(ctx).Authorization(authorization).Action(action).ModifyReleaseRequest(modifyReleaseRequest).Execute()

版本修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyRelease" // string | Action (default to "ModifyRelease")
	modifyReleaseRequest := *openapiclient.NewModifyReleaseRequest("ProjectName_example", int64(123)) // ModifyReleaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyRelease(context.Background()).Authorization(authorization).Action(action).ModifyReleaseRequest(modifyReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRelease`: ModifyRelease200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyRelease&quot;]
 **modifyReleaseRequest** | [**ModifyReleaseRequest**](ModifyReleaseRequest.md) |  | 

### Return type

[**ModifyRelease200Response**](ModifyRelease200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTeamLevelDepotSpec

> ModifyTeamLevelDepotSpec200Response ModifyTeamLevelDepotSpec(ctx).Authorization(authorization).Action(action).ModifyTeamLevelDepotSpecRequest(modifyTeamLevelDepotSpecRequest).Execute()

仓库设置-修改、新增团队级别的仓库规范



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyTeamLevelDepotSpec" // string | Action (default to "ModifyTeamLevelDepotSpec")
	modifyTeamLevelDepotSpecRequest := *openapiclient.NewModifyTeamLevelDepotSpecRequest(*openapiclient.NewDepotSpecTeamLevelParam(false, "Name_example")) // ModifyTeamLevelDepotSpecRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyTeamLevelDepotSpec(context.Background()).Authorization(authorization).Action(action).ModifyTeamLevelDepotSpecRequest(modifyTeamLevelDepotSpecRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyTeamLevelDepotSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTeamLevelDepotSpec`: ModifyTeamLevelDepotSpec200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyTeamLevelDepotSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTeamLevelDepotSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyTeamLevelDepotSpec&quot;]
 **modifyTeamLevelDepotSpecRequest** | [**ModifyTeamLevelDepotSpecRequest**](ModifyTeamLevelDepotSpecRequest.md) |  | 

### Return type

[**ModifyTeamLevelDepotSpec200Response**](ModifyTeamLevelDepotSpec200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTeamMemberLocked

> DeleteAPIDoc200Response ModifyTeamMemberLocked(ctx).Authorization(authorization).Action(action).ModifyTeamMemberUnlockedRequest(modifyTeamMemberUnlockedRequest).Execute()

团队成员锁定



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyTeamMemberLocked" // string | Action (default to "ModifyTeamMemberLocked")
	modifyTeamMemberUnlockedRequest := *openapiclient.NewModifyTeamMemberUnlockedRequest(int32(123)) // ModifyTeamMemberUnlockedRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyTeamMemberLocked(context.Background()).Authorization(authorization).Action(action).ModifyTeamMemberUnlockedRequest(modifyTeamMemberUnlockedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyTeamMemberLocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTeamMemberLocked`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyTeamMemberLocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTeamMemberLockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyTeamMemberLocked&quot;]
 **modifyTeamMemberUnlockedRequest** | [**ModifyTeamMemberUnlockedRequest**](ModifyTeamMemberUnlockedRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTeamMemberUnlocked

> DeleteAPIDoc200Response ModifyTeamMemberUnlocked(ctx).Authorization(authorization).Action(action).ModifyTeamMemberUnlockedRequest(modifyTeamMemberUnlockedRequest).Execute()

团队成员解锁



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyTeamMemberUnlocked" // string | Action (default to "ModifyTeamMemberUnlocked")
	modifyTeamMemberUnlockedRequest := *openapiclient.NewModifyTeamMemberUnlockedRequest(int32(123)) // ModifyTeamMemberUnlockedRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyTeamMemberUnlocked(context.Background()).Authorization(authorization).Action(action).ModifyTeamMemberUnlockedRequest(modifyTeamMemberUnlockedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyTeamMemberUnlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTeamMemberUnlocked`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyTeamMemberUnlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTeamMemberUnlockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyTeamMemberUnlocked&quot;]
 **modifyTeamMemberUnlockedRequest** | [**ModifyTeamMemberUnlockedRequest**](ModifyTeamMemberUnlockedRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTestCase

> ModifyTestCase200Response ModifyTestCase(ctx).Authorization(authorization).Action(action).ModifyTestCaseRequest(modifyTestCaseRequest).Execute()

测试用例修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyTestCase" // string | Action (default to "ModifyTestCase")
	modifyTestCaseRequest := *openapiclient.NewModifyTestCaseRequest(int32(123), "ProjectName_example", int32(123), "TemplateType_example", "Title_example") // ModifyTestCaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyTestCase(context.Background()).Authorization(authorization).Action(action).ModifyTestCaseRequest(modifyTestCaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyTestCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTestCase`: ModifyTestCase200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyTestCase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTestCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyTestCase&quot;]
 **modifyTestCaseRequest** | [**ModifyTestCaseRequest**](ModifyTestCaseRequest.md) |  | 

### Return type

[**ModifyTestCase200Response**](ModifyTestCase200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTestCaseSection

> ModifyTestCaseSection200Response ModifyTestCaseSection(ctx).Authorization(authorization).Action(action).ModifyTestCaseSectionRequest(modifyTestCaseSectionRequest).Execute()

测试用例分组修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyTestCaseSection" // string | Action (default to "ModifyTestCaseSection")
	modifyTestCaseSectionRequest := *openapiclient.NewModifyTestCaseSectionRequest("Name_example", "ProjectName_example", int32(123)) // ModifyTestCaseSectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyTestCaseSection(context.Background()).Authorization(authorization).Action(action).ModifyTestCaseSectionRequest(modifyTestCaseSectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyTestCaseSection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTestCaseSection`: ModifyTestCaseSection200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyTestCaseSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTestCaseSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyTestCaseSection&quot;]
 **modifyTestCaseSectionRequest** | [**ModifyTestCaseSectionRequest**](ModifyTestCaseSectionRequest.md) |  | 

### Return type

[**ModifyTestCaseSection200Response**](ModifyTestCaseSection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTestRun

> ModifyTestRun200Response ModifyTestRun(ctx).Authorization(authorization).Action(action).ModifyTestRunRequest(modifyTestRunRequest).Execute()

测试计划修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyTestRun" // string | Action (default to "ModifyTestRun")
	modifyTestRunRequest := *openapiclient.NewModifyTestRunRequest("ProjectName_example", int32(123)) // ModifyTestRunRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyTestRun(context.Background()).Authorization(authorization).Action(action).ModifyTestRunRequest(modifyTestRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTestRun`: ModifyTestRun200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyTestRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyTestRun&quot;]
 **modifyTestRunRequest** | [**ModifyTestRunRequest**](ModifyTestRunRequest.md) |  | 

### Return type

[**ModifyTestRun200Response**](ModifyTestRun200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWorkItemSplitIssues

> DeleteAPIDoc200Response ModifyWorkItemSplitIssues(ctx).Authorization(authorization).Action(action).ModifyWorkItemSplitIssuesRequest(modifyWorkItemSplitIssuesRequest).Execute()

项目集工作项分解&取消分解到项目中的事项



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyWorkItemSplitIssues" // string | Action (default to "ModifyWorkItemSplitIssues")
	modifyWorkItemSplitIssuesRequest := *openapiclient.NewModifyWorkItemSplitIssuesRequest("IssueCode_example", "ProgramName_example", "ProjectName_example", "Split_example", int64(123)) // ModifyWorkItemSplitIssuesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ModifyWorkItemSplitIssues(context.Background()).Authorization(authorization).Action(action).ModifyWorkItemSplitIssuesRequest(modifyWorkItemSplitIssuesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyWorkItemSplitIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWorkItemSplitIssues`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyWorkItemSplitIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyWorkItemSplitIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyWorkItemSplitIssues&quot;]
 **modifyWorkItemSplitIssuesRequest** | [**ModifyWorkItemSplitIssuesRequest**](ModifyWorkItemSplitIssuesRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanIterationIssue

> DeleteAPIDoc200Response PlanIterationIssue(ctx).Authorization(authorization).Action(action).PlanIterationIssueRequest(planIterationIssueRequest).Execute()

迭代批量规划



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "PlanIterationIssue" // string | Action (default to "PlanIterationIssue")
	planIterationIssueRequest := *openapiclient.NewPlanIterationIssueRequest([]int64{int64(123)}, int64(123), "ProjectName_example") // PlanIterationIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PlanIterationIssue(context.Background()).Authorization(authorization).Action(action).PlanIterationIssueRequest(planIterationIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PlanIterationIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanIterationIssue`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PlanIterationIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanIterationIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;PlanIterationIssue&quot;]
 **planIterationIssueRequest** | [**PlanIterationIssueRequest**](PlanIterationIssueRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseArtifactVersion

> DeleteAPIDoc200Response ReleaseArtifactVersion(ctx).Authorization(authorization).Action(action).DescribeArtifactChecksumsRequest(describeArtifactChecksumsRequest).Execute()

制品版本发布



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ReleaseArtifactVersion" // string | Action (default to "ReleaseArtifactVersion")
	describeArtifactChecksumsRequest := *openapiclient.NewDescribeArtifactChecksumsRequest("Package_example", "PackageVersion_example", int32(123), "Repository_example") // DescribeArtifactChecksumsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ReleaseArtifactVersion(context.Background()).Authorization(authorization).Action(action).DescribeArtifactChecksumsRequest(describeArtifactChecksumsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReleaseArtifactVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseArtifactVersion`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReleaseArtifactVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReleaseArtifactVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ReleaseArtifactVersion&quot;]
 **describeArtifactChecksumsRequest** | [**DescribeArtifactChecksumsRequest**](DescribeArtifactChecksumsRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderCdPipelines

> ReorderCdPipelines200Response ReorderCdPipelines(ctx).Authorization(authorization).Action(action).ReorderCdPipelinesRequest(reorderCdPipelinesRequest).Execute()

部署流程重排序



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ReorderCdPipelines" // string | Action (default to "ReorderCdPipelines")
	reorderCdPipelinesRequest := *openapiclient.NewReorderCdPipelinesRequest("Application_example", []openapiclient.PipelineIdIndex{*openapiclient.NewPipelineIdIndex(int64(123), "PipelineId_example")}) // ReorderCdPipelinesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ReorderCdPipelines(context.Background()).Authorization(authorization).Action(action).ReorderCdPipelinesRequest(reorderCdPipelinesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReorderCdPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderCdPipelines`: ReorderCdPipelines200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReorderCdPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderCdPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ReorderCdPipelines&quot;]
 **reorderCdPipelinesRequest** | [**ReorderCdPipelinesRequest**](ReorderCdPipelinesRequest.md) |  | 

### Return type

[**ReorderCdPipelines200Response**](ReorderCdPipelines200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGrantToResource

> DeleteAPIDoc200Response SetGrantToResource(ctx).Authorization(authorization).Action(action).SetGrantToResourceRequest(setGrantToResourceRequest).Execute()

授权设置，收回授权体在资源中的所有授权，重新设置为参数指定的授权



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "SetGrantToResource" // string | Action (default to "SetGrantToResource")
	setGrantToResourceRequest := *openapiclient.NewSetGrantToResourceRequest([]openapiclient.GrantInfo{*openapiclient.NewGrantInfo("GrantObjectId_example", "GrantScope_example", int64(123))}, *openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example")) // SetGrantToResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SetGrantToResource(context.Background()).Authorization(authorization).Action(action).SetGrantToResourceRequest(setGrantToResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetGrantToResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGrantToResource`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetGrantToResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGrantToResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;SetGrantToResource&quot;]
 **setGrantToResourceRequest** | [**SetGrantToResourceRequest**](SetGrantToResourceRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPredicatePolicyOnResource

> DeleteAPIDoc200Response SetPredicatePolicyOnResource(ctx).Authorization(authorization).Action(action).SetPredicatePolicyOnResourceRequest(setPredicatePolicyOnResourceRequest).Execute()

资源权限判定策略设置



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "SetPredicatePolicyOnResource" // string | Action (default to "SetPredicatePolicyOnResource")
	setPredicatePolicyOnResourceRequest := *openapiclient.NewSetPredicatePolicyOnResourceRequest(*openapiclient.NewResourceInfo("ResourceId_example", "ResourceType_example"), "ResourcePredicatePolicy_example") // SetPredicatePolicyOnResourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SetPredicatePolicyOnResource(context.Background()).Authorization(authorization).Action(action).SetPredicatePolicyOnResourceRequest(setPredicatePolicyOnResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetPredicatePolicyOnResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPredicatePolicyOnResource`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetPredicatePolicyOnResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPredicatePolicyOnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;SetPredicatePolicyOnResource&quot;]
 **setPredicatePolicyOnResourceRequest** | [**SetPredicatePolicyOnResourceRequest**](SetPredicatePolicyOnResourceRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCodingCIBuild

> DeleteAPIDoc200Response StopCodingCIBuild(ctx).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()

构建停止



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "StopCodingCIBuild" // string | Action (default to "StopCodingCIBuild")
	describeCodingCIBuildStageRequest := *openapiclient.NewDescribeCodingCIBuildStageRequest(int32(123)) // DescribeCodingCIBuildStageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.StopCodingCIBuild(context.Background()).Authorization(authorization).Action(action).DescribeCodingCIBuildStageRequest(describeCodingCIBuildStageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StopCodingCIBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCodingCIBuild`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StopCodingCIBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopCodingCIBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;StopCodingCIBuild&quot;]
 **describeCodingCIBuildStageRequest** | [**DescribeCodingCIBuildStageRequest**](DescribeCodingCIBuildStageRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerCdPipeline

> TriggerCdPipeline200Response TriggerCdPipeline(ctx).Authorization(authorization).Action(action).TriggerCdPipelineRequest(triggerCdPipelineRequest).Execute()

部署流程触发



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "TriggerCdPipeline" // string | Action (default to "TriggerCdPipeline")
	triggerCdPipelineRequest := *openapiclient.NewTriggerCdPipelineRequest("Application_example", "PipelineNameOrId_example", "TriggerJsonContent_example") // TriggerCdPipelineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.TriggerCdPipeline(context.Background()).Authorization(authorization).Action(action).TriggerCdPipelineRequest(triggerCdPipelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TriggerCdPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerCdPipeline`: TriggerCdPipeline200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TriggerCdPipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerCdPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;TriggerCdPipeline&quot;]
 **triggerCdPipelineRequest** | [**TriggerCdPipelineRequest**](TriggerCdPipelineRequest.md) |  | 

### Return type

[**TriggerCdPipeline200Response**](TriggerCdPipeline200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerCodingCIBuild

> TriggerCodingCIBuild200Response TriggerCodingCIBuild(ctx).Authorization(authorization).Action(action).TriggerCodingCIBuildRequest(triggerCodingCIBuildRequest).Execute()

构建触发



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "TriggerCodingCIBuild" // string | Action (default to "TriggerCodingCIBuild")
	triggerCodingCIBuildRequest := *openapiclient.NewTriggerCodingCIBuildRequest(int32(123)) // TriggerCodingCIBuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.TriggerCodingCIBuild(context.Background()).Authorization(authorization).Action(action).TriggerCodingCIBuildRequest(triggerCodingCIBuildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TriggerCodingCIBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerCodingCIBuild`: TriggerCodingCIBuild200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TriggerCodingCIBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerCodingCIBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;TriggerCodingCIBuild&quot;]
 **triggerCodingCIBuildRequest** | [**TriggerCodingCIBuildRequest**](TriggerCodingCIBuildRequest.md) |  | 

### Return type

[**TriggerCodingCIBuild200Response**](TriggerCodingCIBuild200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerCodingScan

> TriggerCodingScan200Response TriggerCodingScan(ctx).Authorization(authorization).Action(action).TriggerCodingScanRequest(triggerCodingScanRequest).Execute()

代码扫描触发



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "TriggerCodingScan" // string | Action (default to "TriggerCodingScan")
	triggerCodingScanRequest := *openapiclient.NewTriggerCodingScanRequest(int32(123)) // TriggerCodingScanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.TriggerCodingScan(context.Background()).Authorization(authorization).Action(action).TriggerCodingScanRequest(triggerCodingScanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TriggerCodingScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerCodingScan`: TriggerCodingScan200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TriggerCodingScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerCodingScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;TriggerCodingScan&quot;]
 **triggerCodingScanRequest** | [**TriggerCodingScanRequest**](TriggerCodingScanRequest.md) |  | 

### Return type

[**TriggerCodingScan200Response**](TriggerCodingScan200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserGroupById

> DeleteAPIDoc200Response UpdateUserGroupById(ctx).Authorization(authorization).Action(action).UpdateUserGroupByIdRequest(updateUserGroupByIdRequest).Execute()

用户组更新



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "UpdateUserGroupById" // string | Action (default to "UpdateUserGroupById")
	updateUserGroupByIdRequest := *openapiclient.NewUpdateUserGroupByIdRequest("Description_example", int64(123), "Name_example") // UpdateUserGroupByIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.UpdateUserGroupById(context.Background()).Authorization(authorization).Action(action).UpdateUserGroupByIdRequest(updateUserGroupByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserGroupById`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserGroupById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;UpdateUserGroupById&quot;]
 **updateUserGroupByIdRequest** | [**UpdateUserGroupByIdRequest**](UpdateUserGroupByIdRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

