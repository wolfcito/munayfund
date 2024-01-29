# Munayfund

![list](/public/app-screenshots/home.png)

Our crowdfunding platform, powered by Arbitrum and Allo Protocol, redefines the
decentralized funding experience. It enables users to create and fund projects
using cryptocurrencies on the Arbitrum network. With a common fund backed by
Allo Protocol, participants vote to decide the fund distribution among projects.
Additionally, it incorporates machine learning for project evaluation, a
user-friendly interface, and robust security in smart contracts, ensuring
transparency and active community participation.

## Architecture
MunayFund uses the following tech:

- Solidity
- Reactjs (Nextjs)
- Allo protocol + arbitrum
- IPFS
- Go
- Docker

## Features
- Create Jars of Funds (commonly known as Pools)
- Donate or transfer funds to jars.
- Crate projects and apply to active jars.
- Vote on projects (Allocate).
- Submit proofs of deliveries, photos, or links to the work done.
- Approve proofs of deliveries and allow projects to recieve the funds from the jar.
- Distribute rewards to projects based on votes, time interval and deliveries approvals.

### The Cookie Jar Strategy
We manage a logic of founding of two sides, one via crowfounding and other with a Cookie jar. The cookie jar is the result of a single found that is filled by other users and at the end of the period the founding is shared on different percentages between the multiple projects that have advances and are signaled as trusted projects.
This allows for verification of work before distributing the funds, thus filtering out abandoned or scam projects.

## App Screenshots
### Home Page

![home](/public/app-screenshots/home.png)
![about](/public/app-screenshots/home-2.png)

### Create Allo Profiles

![profile](/public/app-screenshots/profile.png)

### Admin Dashboard (Jar Creator)

![dashboard](/public/app-screenshots/dashboard.png)

### View Jars

![view-jars](/public/app-screenshots/jars.png)
![jar-details](/public/app-screenshots/jardetails.png)

### Create Jars

![new-jar](/public/app-screenshots/newjar.png)
![new-jar-2](/public/app-screenshots/newjar-2.png)

### Vote on projects

![project-details](/public/app-screenshots/projectdetails.png)

## Resources
- [Allo Protocol Docs](https://docs.allo.gitcoin.co/)
- [Allo Protocol Repo](https://github.com/allo-protocol/allo-v2)

## License
MIT. This project is open source and open to collaboration.