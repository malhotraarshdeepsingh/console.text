import {Pulsey} from 'pulsey';

const pulsey = Pulsey('your-api-key-here');

const text = await pulsey.text('500', 'Alert: Something broke.');
console.log('Pulsey Text Response:', text);

const siren = await pulsey.siren('500', 'Critical: Call the dev team now!');
console.log('Pulsey Siren Response:', siren);
