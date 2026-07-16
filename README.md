# hr-auction-hlf

Solution repository for the **Open Auction** Hyperledger Fabric chaincode
challenge (NPCI / HackerRank, Hard).

Standard Fabric **test-network** plus a chaincode skeleton at
[`chaincode/auction.go`](chaincode/auction.go). Cloned into the candidate's
environment by the HackerRank Setup Script (via [`setup.sh`](setup.sh)).

## Candidate task
1. Implement the six functions in `chaincode/auction.go`, including the
   ascending-bid rule and winner determination.
2. Deploy: `cd test-network && ./network.sh deployCC -ccn auctioncc -ccp ../chaincode -ccl go`
3. auc1: create, bid Alice 500, bid Bob 750, close (winner Bob). auc2: create, cancel.

---

Authored by **Dayal Mukati** — [dayalmukati.com](https://dayalmukati.com)
