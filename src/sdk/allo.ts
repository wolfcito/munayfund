import { getIPFSClient } from '@/services/ipfs'
import { Allo } from '@allo-team/allo-v2-sdk'
import { deployMicrograntsStrategy } from './microgrants'

// create an allo instance
// todo: snippet => createAlloInstance
export const allo = new Allo({
  chain: 421614,
  rpc: 'https://arbitrum-sepolia.blockpi.network/v1/rpc/public',
})

export const createPool = async () => {
  // Create a profile to use as the pool owner/creator
  // todo: you can add the profileId you generated intiially here so you don't have to create a new one each time.
  const profileId = '0xbbc8d31c2b00a011912322740d238139ae8578a4dec18818408f720bdcb54b53'
  // const profileId = await createProfile();

  // Save metadata to IPFS -> returns a pointer we save on chain for the metadata
  const ipfsClient = getIPFSClient()
  const metadata = {
    profileId: profileId,
    name: 'MunayFund',
    website: 'https://munayfund.vercel.app',
    description: 'The importance of support that goes beyond the material aspect',
    base64Image: '', // skipping the image for the demo
  }

  // NOTE: Use this to pin your base64 image to IPFS
  // let imagePointer;
  // if (metadata.base64Image && metadata.base64Image.includes("base64")) {
  //   imagePointer = await ipfsClient.pinJSON({
  //     data: metadata.base64Image,
  //   });
  //   metadata.base64Image = imagePointer;
  // }

  const pointer = await ipfsClient.pinJSON(metadata)
  console.log('Metadata saved to IPFS with pointer: ', pointer)

  // Deploy the microgrants strategy - `microgrants.ts`
  const poolId = await deployMicrograntsStrategy(pointer, profileId)
  console.log('Pool created with ID: ', poolId)

  return poolId
}
