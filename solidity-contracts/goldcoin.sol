// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// https://docs.soliditylang.org/en/v0.8.15/path-resolution.html#base-path-and-include-paths to compile need to explicitly include nodemodules path while compiling
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Goldcoin is ERC20 {
    constructor() ERC20("Goldcoin", "GLD") {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
}