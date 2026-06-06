import { FeatureId } from '@/react/portainer/feature-flags/enums';
import { isLimitedToBE } from '@/react/portainer/feature-flags/feature-flags.service';

const BE_URL = 'https://opendocking.io/';

export function getFeatureDetails(featureId?: FeatureId) {
  if (!featureId) {
    return {};
  }
  const url = `${BE_URL}${featureId}`;

  const limitedToBE = isLimitedToBE(featureId);

  return { url, limitedToBE };
}
