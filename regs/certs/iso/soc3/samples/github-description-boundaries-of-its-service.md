
The following was taken from GitHub Inc.'s SOC 3 report on 12th August 2024. This description will be useful when writing your own.

# GitHub, Inc.'s Description of the Boundaries of Its GitHub Enterprise Cloud service

# Type of Services Provided

GitHub (“the Company”) is an independently operated subsidiary of Microsoft and generated its first commit in 2007. It’s headquartered in San Francisco, California, with additional offices in Bellevue, WA and Oxford, UK. GitHub currently employs approximately 3,000 employees, with approximately 95 percent of the workforce being remote.
GitHub is a web-based software development platform built on the Git version control software. Primarily used for software code, GitHub offers the distributed version control and source code management functionality of Git with additional features and enhancements. Specifically, it provides access control and several collaboration features including bug tracking, feature requests, task management, GitHub Advanced Security, GitHub Actions, GitHub Teams, pull requests, discussions, issues, pages, projects, docs, and wikis.
GitHub Enterprise Cloud service is GitHub’s SaaS solution for collaborative software development. Features of the Enterprise Cloud service include:

## Organizations

An organization is a collection of user accounts that owns repositories. Organizations have one or more owners, who have administrative privileges for the organization. When a user creates an organization, it does not have any repositories associated with it. At any time, members of the organization with the Owner role can add new repositories or transfer existing repositories.

## Code Hosting

GitHub is one of the largest code hosts in the world, with millions of projects. Private, public, or open source repositories are equipped with tools to host, version, and release code. Unlimited private repositories allow keeping the code in one place, even when using Subversion (SVN) or working with large files using Git Large File Storage (LFS).
Changes can be made to code in precise commits allowing for quick searches on commit messages in the revision history to find a change. In addition, blame view enables users to trace changes and discover how the file and code base has evolved.
With sharing, changes can be packaged from a recently closed milestone or finished project into a new release. Users can draft and publish release notes, publish pre-release versions, attach files, and link directly to the latest download.

## Code Management

Code review is a critical path to better code, and it’s fundamental to how GitHub works. Built-in review tools make code review an essential part of team development workflows.
A pull request (PR) is a living conversation where ideas can be shared, tasks assigned, details discussed, and reviews conducted. Reviews happen faster when GitHub shows a user exactly what has changed. Diffs compare versions of source code side by side, highlighting the parts that are new, edited, or deleted.
PRs also enable clear feedback, review requests, and comments in context with comment threads within the code. Comments may be bundled into one review or in reply to someone else’s comments inline as a conversation.

## Project Management

Project boards allow users to reference every issue and PR in a card, providing a drag-and-droppable snapshot of the work that teams do in a repository. This feature can also function as an agile idea board to capture early ideas that come up as part of a standup or team sync, without polluting the issues.
Issues enable team task tracking, with resources identified and assigned tasks within a team. Issues may be used to track a bug, discuss an idea with an @mention, or start distributing work. Issue and PR assignments to one or more teammates make it clear who is doing what work and what feedback and approvals have been requested.
Milestones can be added to issues or PRs to organize and track progress on groups of issues or PRs in a repository.

## Team Management

Building software is as much about managing teams and communities as it is about code. Users set roles and expectations without starting from scratch. Customized common codes of conduct can be created for any project, with pre-written licenses available right from the repository.
GitHub Teams organizes people, provides level-up access with administrative roles, and tunes permissions for nested teams. Discussion threads keep conversations on topic using moderation tools, like issue and pull-request locking, to help teams stay focused on code. For maintaining open source projects, user blocking reduces noise and keeps conversations productive.

## Documentation

GitHub allows for documentation to be created and maintained in any repository, and wikis are available to create documentation with version control. Each wiki is its own repository, so every change is versioned and comparable. With a text editor, users can add docs in the text formatting language of choice, such as Textile or GitHub Flavored Markdown.

## GitHub Actions

GitHub Actions automates Continuous Integration/ Continuous Delivery (CI/CD) software workflows by enabling the build, test, and deployment of code directly from GitHub, with code reviews, branch management, and issue triaging customized to work the way developers need.
GitHub Actions initiates workflows for events like push, issue creation, or a new release, and actions can be combined and configured for the services used, built, and maintained by the community.
GitHub Actions supports additional options to do things like build containers, deploy web services, or automate notifications to users of open source projects using GitHub developers’ existing GITHUB_TOKEN in collaboration with other Enterprise Cloud features.

## GitHub Advanced Security

GitHub Advanced Security provides features that help improve and maintain the quality and security of code. Code scanning searches for potential security vulnerabilities and coding errors in code using CodeQL or a third-party tool. Secret scanning detects secrets, for example keys and tokens, that have been inadvertently checked into repositories. If push protection is enabled, secret scanning also detects secrets and blocks contributors from pushing them to repositories. Dependency review and Dependabot detect vulnerable versions of dependencies and warn about the associated security vulnerabilities. Alerts from all features can be centrally reported and tracked.

## System Boundaries

The scope of this report includes GitHub’s Enterprise Cloud, and the supporting production systems, infrastructure, software, people, procedures, and data. The following Enterprise Cloud features are included in the scope of this report: Issues, PRs, Discussions, Wikis, Pages, Projects, Docs, Audit Logging, GitHub Advanced Security, GitHub Teams, Dependabot, and GitHub Actions.

## The Components of the System Used to Provide the Services

The boundaries of GitHub Enterprise Cloud service are the specific aspects of the Company’s infrastructure, software, people, procedures, and data necessary to provide its services and that directly support the services provided to customers. Any infrastructure, software, people, procedures, and data that indirectly support the services provided to customers are not included within the boundaries of GitHub Enterprise Cloud service.

### Infrastructure

The Company utilizes a third party cloud service provider and GitHub-managed colocation data centers to provide the resources to host GitHub Enterprise Cloud service. The Company leverages the experience and resources of the third party cloud service provider to scale quickly and securely as necessary to meet current and future demand. However, the Company is responsible for designing and configuring the GitHub Enterprise Cloud service architecture within the third party cloud service provider and GitHub-managed colocation data centers to ensure security and resiliency requirements are met.
The in-scope hosted infrastructure also consists of multiple supporting tools.

### Software

Software consists of the programs and software that support GitHub Enterprise Cloud service (operating systems [OSs], middleware, and utilities). The list of software and ancillary software used to build, support, secure, maintain, and monitor GitHub Enterprise Cloud service include the following applications:
- CI/CD
- Dependency Management, Static Code Analysis, Secret Scanning
- Network architecture configuration and management
- Application and infrastructure monitoring
- Endpoint management and security
- In-house developed tool for data center asset inventory and management
- Account and license storage
- Single Sign-On (SSO) configuration management
- Infrastructure Configuration Management
- Cloud asset inventory
- ChatOps and daily work communications
- Security information and event management (SIEM), logging system, intrusion detection
- Vulnerability scanning
- OS Baseline
- Customer Support

### People

The Company develops, manages, and secures GitHub Enterprise Cloud service via separate departments. The responsibilities of these departments are defined in the following table:

| **Group/Role Name** | **Function** |
|--------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| Customer Support | Responsible for providing technical and account-related support to GitHub Enterprise Cloud service customers and for resolving customer issues via email, chat, social media, and phone from developers and customer entities around the globe. |
| Engineering | Responsible for working with the Product team to plan and coordinate releases, and accountable for building, testing, and deploying GitHub Enterprise Cloud service code and feature changes.<br>Responsible for maintaining service availability, including performance and scale monitoring and reporting, incident command, and on-call readiness for any production issues. Responsible for configuration management, building, testing, and deploying software relevant to the operation and management of production assets, patching and remediation of vulnerabilities reported by the Security team, and data center operations management. Responsible for managing Git and database storage backups and restores. |
| Legal | Responsible for negotiating contractual obligations with third parties and technology partners/suppliers, legal terms and conditions, and ensuring compliance with internal contractual standards. |
| People Operations | Responsible for talent acquisition, diversity and inclusion, learning and development, and employee engagement on everything from benefits and perks to career development and growth. |
| Privacy | Responsible for determining which privacy laws and regulations apply to GitHub and determining the best way to comply with them, ultimately ensuring GitHub can offer its products to every developer anywhere in the world. |
| Product Management | Responsible for understanding customer requirements, collecting, defining, and clarifying feature requests and development efforts, and managing feature rollouts and related customer communication efforts. |
| Security | Responsible for ensuring the security of GitHub products. Security consists of multiple teams with specific missions: Threat Hunting Operations and Response Incident Response Team (THORIR), Product Security Incident Response Team (PSIRT), Security Lab, Security Operations, Secure Access Engineering, Security Telemetry, Vulnerability Management, GitHub Cloud and Enterprise Security, and Governance, Risk, Compliance, and Communication (GRCC). These teams manage security incident detection and response, monitoring, vulnerability scanning, network and application layer penetration testing, security architecture, security engineering and operations, access management, endpoint asset management, and risk and compliance oversight. |
| Senior Leadership | Responsible for the overall governance of GitHub. This group includes the Chief Executive Officer (CEO), Head of Finance, Chief Revenue Officer (CRO), Chief Security Officer (CSO), Chief Human Resource Officer (CHRO), Head of Design, Chief Operations Officer, Chief of Staff, Vice President, Senior Vice President of Engineering, Chief Legal Officer, and Vice President of Communities. |
---
## Policies, Standards, and Procedures
GitHub maintains Policies, Standards, and Procedures necessary to securely operate GitHub Enterprise Cloud service. Policies and Standards are centrally managed in The Hub, GitHub’s centralized internal communication platform. The Hub is backed by a repository, which is used to implement annual reviews and control changes to Policies and Standards. Once a change has been approved by the owner of the Policy or Standard, it is automatically updated on The Hub. Procedures are developed and documented within the GitHub repositories maintained by every team to provide end-user documentation and guidance on the multitude of operational functions performed daily by GitHub security and product engineers, developers, administrators, and support personnel. These procedures are drafted in alignment with the overall Policies and Standards and are updated as necessary to reflect changes in the business.
GitHub Policies and Standards establish controls to enable security, efficiency, availability, and quality of service. The GitHub Information Security and Privacy Management System (ISPMS) Policy and related policies define information security practices, roles, and responsibilities. The ISPMS outlines the security roles and responsibilities for the organization and expectations for employees, contractors, and third parties utilizing GitHub systems or data.
This overarching security policy is supported by several dependent security policies, standards, and procedures applicable to the operation and management of security across the organization. Security-related policies, standards, and procedures are documented and made available to individuals responsible for their implementation and compliance.
---
### Below is the current inventory of security and audit-related policies and standards that inform procedures operating in support of the GitHub ISPMS Policy objectives:
| **Policy** | **Associated Standards and Procedures** |
|----------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| GitHub ISPMS | No Associated Standards or Procedures |
| GitHub ISPMS Scope | No Associated Standards or Procedures |
| GitHub ISPMS Statement of Applicability (SOA) | • ChatOps Command Security and Risk Standard<br>• Controls Monitoring Standard<br>• Controls Monitoring SOP<br>• Data Classification Standard<br>• Domain Management Standard<br>• Endpoint Security Standard<br>• Enterprise Administration Standard<br>• External File Sharing Standard<br>• Git Systems Server Site Failure plan<br>• High-Risk Application Access Standard<br>• Organization Administration Standard<br>• Production VPN Access Standard<br>• Repository Security Baseline Configuration Standard<br>• Reviewing Pull Requests<br>• Server Operating System Standard |
| Corporate Data Retention Policy | • Audit Video Retention Standard<br>• Corporate Data Retention Standard<br>• Product Telemetry Data Retention Standard<br>• Slack Retention Standard |
| Contractor Termination Policy | No Associated Standards or Procedures |
| Full-Time Employee Termination Policy | No Associated Standards or Procedures |
| Identity and Access Management Policy | • Identity and Access Management Standard<br>• IAM Onboarding SOP<br>• IAM Entitlements SOP<br>• IAM Privileged Systems and Elevated Access SOP<br>• Granting Slack Access to Contractors and Consultants SOP<br>• IAM Non-Human Accounts in Okta<br>• IAM Offboarding SOP<br>• IAM On-Leave SOP |
| Physical and Environmental Protection Policy | • Production Datacenter Standard<br>• Datacenter Physical Access SOP<br>• Production Media Destruction SOP<br>• Datacenter Access compliance guidelines |
| Privacy Statement | No Associated Standards or Procedures |
| Private Information Removal Policy - External Customer Facing Policy | No Associated Standards or Procedures |

### Data

Data refers to transaction streams, files, data stores, tables, and output used or processed by the Company. GitHub uses repository data to connect users to relevant tools, people, projects, and information. Repositories are categorized as either public, private, or open source. Public repositories can be viewed by anyone, including people who are not GitHub users. Private repositories are only visible to the repository owner and collaborators that the owner specified. GitHub aggregates metadata and parses content patterns to deliver generalized insights within the product. It uses data from public repositories, and uses metadata and aggregate data from private repositories when a repository's owner has chosen to share the data with GitHub through an opt-in.
If a private repository is opted in for data use to take advantage of any of the capabilities of the security and analysis features, then GitHub will perform a read-only analysis of that specific private repository's git contents. If a private repository is not opted in for data use, its private data, source code, or trade secrets are classified internally as restricted, and they are maintained as confidential and private consistent with GitHub's Terms of Service.
Customer data is managed, processed, and stored in accordance with relevant data protection and other regulations and with specific requirements formally established in client contracts.
The Company has deployed secure methods and protocols for transmission of confidential or sensitive information over public networks.

## Subservice Organizations

The Company uses subservice organizations for data center colocation services and infrastructure hosting. The Company’s controls related to GitHub Enterprise Cloud service cover only a portion of the overall internal control for each user entity of GitHub Enterprise Cloud service. The description does not extend to the colocation services for IT infrastructure provided by the subservice organizations.
Although the subservice organizations have been carved out for the purposes of this report, certain service commitments, system requirements, and applicable criteria are intended to be met by controls at the subservice organizations. Controls are expected to be in place at the subservice organizations related to physical security and environmental protection. The subservice organizations’ physical security controls should mitigate the risk of unauthorized access to the hosting facilities. The subservice organizations’ environmental protection controls should mitigate the risk of fires, power loss, climate, and temperature variabilities.
The Company management receives and reviews the subservice organization’s SOC 2 reports, International Standards Organization (ISO) 27001 and 27017 Certifications, and Payment Card Industry Attestation of Compliance (PCI AOC) as they are issued, and at least annually. In addition, through its operational activities, Company management monitors the services performed by the subservice organizations to determine whether operations and controls expected to be implemented are functioning effectively. Management also communicates with the subservice organizations to monitor compliance with the service agreement, stay informed of changes planned at the hosting facility, and relay any issues or concerns to management of the subservice organizations.
