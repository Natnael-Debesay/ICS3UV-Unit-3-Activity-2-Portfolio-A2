/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-11-18
 * @fileoverview This program displays the ten Canadian provinces, three territories, and their capital cities.
 */

// variables
let ProvinceTerritory;
let Capital; 

// input
ProvinceTerritory = prompt("Province/Territory") || "No province/territory entered!";
Capital = prompt("Capital") || "-1";

// output
console.log("\n");
console.log(ProvinceTerritory + ": " + Capital);

console.log("/nDone.");
