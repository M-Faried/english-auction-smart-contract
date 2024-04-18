// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

interface IERC721 {
    function transferFrom(address from, address to, uint tokenId) external;
}

contract EnglishAuction {
    // Auction states
    bool public started;
    bool public ended;
    uint public endTime;
    uint public highestBid;
    address public highestBidder;
    mapping(address => uint) public allBids;

    // For constructor
    address payable public immutable owner;
    IERC721 public immutable nft;
    uint public immutable nftId;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call the function");
        _;
    }

    modifier requireStarted() {
        require(started, "Auction has not started");
        require(!ended, "Auction has ended");
        _;
    }

    event Started(uint startTime, uint endTime);
    event Bid(address indexed bidder, uint value);
    event Ended(address indexed highestBidder, uint value);

    constructor(address _nft, uint _nftId) {
        owner = payable(msg.sender);
        nft = IERC721(_nft);
        nftId = _nftId;
    }

    function bid() external payable requireStarted {
        require(
            msg.value > highestBid,
            "Bid price is lower than the current highest bid"
        );
        require(block.timestamp < endTime, "Auction has ended");

        highestBidder = msg.sender;
        highestBid = msg.value;
        allBids[msg.sender] = msg.value;

        emit Bid(msg.sender, msg.value);
    }

    function start(uint _openingBid, uint _duration) external onlyOwner {
        require(!started, "Auction has already started");

        highestBid = _openingBid;
        endTime = block.timestamp + _duration;
        nft.transferFrom(owner, address(this), nftId);
        started = true;
        emit Started(block.timestamp, endTime);
    }

    function end() external onlyOwner requireStarted {
        require(block.timestamp >= endTime, "Auction has ended");

        ended = true;
        nft.transferFrom(address(this), highestBidder, nftId);
        owner.transfer(highestBid);

        emit Ended(highestBidder, highestBid);
    }

    function withdraw() external {
        // bidder to receive fund from all the bids state.
    }
}
