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

    event Started(uint startTime, uint endTime);

    constructor(address _nft, uint _nftId){
        owner = payable(msg.sender);
        nft = IERC721(_nft);
        nftId = _nftId;
    }

    function bid() external payable{
        // validations    
        // highest bidder, all bids
    }

    function start(uint _openingBid, uint _duration) external onlyOwner {

        require(!started, "Auction has already started");        

        highestBid = _openingBid;
        endTime = block.timestamp + _duration;
        nft.transferFrom(owner, address(this), nftId);
        started = true;
        emit Started(block.timestamp, endTime);
    }

    function end() external onlyOwner {
        // highest bidder receive the itme
        // owner recieves ether
    }

    function withdraw() external {
        // bidder to receive fund from all the bids state.
    }
    
}
